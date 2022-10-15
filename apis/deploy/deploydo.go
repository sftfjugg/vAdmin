// Copyright 2019 syncd Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package deploy

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"time"

	"vAdmin/module/deploy"
	"vAdmin/render"
	"vAdmin/tools/gostring"
	"vAdmin/pkg/logger"
	"vAdmin/tools/ws"
	"fmt"
)

func DeployStop(c *gin.Context) {
	id := gostring.Str2Int(c.PostForm("id"))
	if id == 0 {
		render.ParamError(c, "id cannot be empty")
		return
	}
	apply := &deploy.Apply{
		ID: id,
	}
	if err := apply.Detail(); err != nil {
		render.AppError(c, err.Error())
		return
	}

	if !ExistsTask(id) {
		apply := &deploy.Apply{
			ID: id,
			Status: deploy.APPLY_STATUS_FAILED,
		}
		apply.UpdateStatus()
	} else {
		StopTask(id)
	}
	render.JSON(c, nil)
}

func DeployStatus(c *gin.Context) {
	id := gostring.Str2Int(c.Query("id"))
	if id == 0 {
		render.ParamError(c, "id cannot be empty")
		return
	}

	apply := &deploy.Apply{
		ID: id,
	}
	if err := apply.Detail(); err != nil {
		render.AppError(c, err.Error())
		return
	}

	taskRest := []map[string]interface{}{}
	depResults := StatusTask(id)
	if depResults != nil  {
		for _, depResult := range depResults {
			groupStatus := deploy.DEPLOY_STATUS_NONE
			switch depResult.Status {
			case STATUS_INIT:
				groupStatus = deploy.DEPLOY_STATUS_NONE
			case STATUS_ING:
				groupStatus = deploy.DEPLOY_STATUS_START
			case STATUS_DONE:
				groupStatus = deploy.DEPLOY_STATUS_SUCCESS
			case STATUS_FAILED:
				groupStatus = deploy.DEPLOY_STATUS_FAILED
			}

			srvRest := []map[string]interface{}{}
			for _, r := range depResult.ServerRest {
				var err string
				if e := r.Error; e != nil {
					err = e.Error()
				}
				srvRest = append(srvRest, map[string]interface{}{
					"id": r.ID,
					"task": r.TaskResult,
					"status": r.Status,
					"error": err,
				})
			}
			groupRest := map[string]interface{}{
				"group_id": depResult.ID,
				"status": groupStatus,
				"content": srvRest,
			}
			taskRest = append(taskRest, groupRest)
		}
	} else {
		d := &deploy.Deploy{
			ApplyId: id,
		}
		taskList, err := d.TaskList()
		if err != nil {
			render.AppError(c, err.Error())
			return
		}
		for _, l := range taskList {
			var obj interface{}
			gostring.JsonDecode([]byte(l.Content), &obj)

			groupRest := map[string]interface{}{
				//"group_id": l.GroupId,
				"status": l.Status,
				"content": obj,
			}
			taskRest = append(taskRest, groupRest)
		}
	}
	render.JSON(c, map[string]interface{}{
		"status": apply.Status,
		"task_list": taskRest,
	})
}

func DeployStart(c *gin.Context) {

	/*
	var data models.DeployApp

	// ShouldBindJSON解析JSON数据，但在重复调用的情况下会出现EOF的报错
	err := c.ShouldBindJSON(&data)
	id := data.ID
	//id := gostring.Str2Int(c.PostForm("id"))
	fmt.Printf("id------------------>: %d\n",id)

	//deploy_mode如果存到了其他表，需要单独在这里获取
	//isParallel := c.PostForm("deploy_mode")
	//fmt.Printf("isParallel------------------>: %s\n",isParallel)

	// 判断上线单ID是否有效
	if id == 0 {
		render.ParamError(c, "apply id cannot be empty")
		return
	}

	// 上线单实例化
	apply := &deploy.Apply{
		ID: id,
	}

	// 判断上线单详情
	if err := apply.Detail(); err != nil {
		render.AppError(c, err.Error())
		return
	}

	// 获取部署模式：并行还是串行
	var isParallel string
	if apply.DeployMode == "" {
		render.AppError(c, "deploy apply can not get deploy_mode")
		return
	} else {
		isParallel = apply.DeployMode
		fmt.Printf("isParallel------------------>: %s\n",isParallel)
	}

	// 部署前状态检查：只能部署状态是没上线或者是上线失败的
	if apply.Status != deploy.APPLY_STATUS_NONE && apply.Status != deploy.APPLY_STATUS_FAILED {
		render.AppError(c, "deploy apply have deployed success")
		return
	}

	// 部署前事务检查：不能部署状态是正在上线中的
	if canDeploy, err := apply.CheckHaveDeploying(); !canDeploy || err != nil {
		if err != nil {
			fmt.Println(err.Error())
		}
		render.RepeatError(c, "project have deploying apply within 24 hours")
		return
	}

	// 上线单必须审核通过
	if apply.AuditStatus != deploy.AUDIT_STATUS_OK {
		fmt.Println("apply audit_status must passed")
	}
	 */

	var (
		wsConn *websocket.Conn
		err    error
		//data   []byte
	)

	// get host once before establish websocket connection,
	// if error happened, do not establish websocket connection.
	// Upgrade: websocket
	wsConn, err = upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Error(err)
		//EmptyHostResponse(c, fmt.Sprintf("获取主机[%s]的websocket连接建立失败", uuid))
		return
	}
	conn := ws.InitConnection(wsConn)
	logger.Info(fmt.Sprintf("depolying [%s] websocket success", "w"))


	go func(conn *ws.Connection, Dversion string) {
		for {
			//Did := []byte(DeployInfo.DemandID)
			err = conn.WriteMessage([]byte(Dversion))
			if err != nil {
				conn.Close()
				return
			}
			time.Sleep(time.Second * time.Duration(wsUpdateInterval))
		}
	}(conn, "w")
	/*
	apply.Status = deploy.APPLY_STATUS_ING
	if err := apply.UpdateStatus(); err != nil {
		fmt.Println(err.Error())
	}



	fmt.Println("VVVVVVVVVVVVVVVVVVVVVVVVVVVVVV---->y2")
	go func() {
		a := <-ExecStatusChan // 没有记录的话会堵塞，会等下次的记录被插入才会放行
		fmt.Printf("VVVVVVVVVVVVVVVVVVVVVVVVVVVVVV---->%d",a)
		if a == STATUS_FAILED {
			apply.Status = deploy.APPLY_STATUS_FAILED
			if err := apply.UpdateStatus(); err != nil {
				fmt.Println(err.Error())
			}
			//deployStatus = MAIL_STATUS_FAILED
		} else if a == STATUS_DONE {
			apply.Status = deploy.APPLY_STATUS_SUCCESS
			if err := apply.UpdateStatus(); err != nil {
				fmt.Println(err.Error())
			}
			//deployStatus = MAIL_STATUS_SUCCESS
		}
	}()

	*/
	// Write task init to DB
	/*
	d := &deploy.Deploy{
            ApplyId: id,
            GroupId: gid,
            Status:  deploy.DEPLOY_STATUS_NONE,
        }
        if err := d.Create(); err != nil {
            render.AppError(c, err.Error())
            return
        }
	*/

	//send deploy email
	/*
		if err := apply.Detail(); err != nil {
			return
		}
		u := &user.User{
			ID: apply.UserId,
		}
		if err := u.Detail(); err != nil {
			return
		}
		proj := &project.Project{
			ID: apply.ProjectId,
		}
		if err := proj.Detail(); err != nil {
			return
		}
		mails := gostring.JoinStrings(proj.DeployNotice, ",", u.Email)
		MailSend(&MailMessage{
			Mail: mails,
			ApplyId: id,
			Mode: MAIL_MODE_DEPLOY,
			Status: deployStatus,
			Title: apply.Name,
		})
	*/
	// run hook script
	// common.HookDeploy(id)
	//render.JSON(c, nil)
}

/*
func KillTask(c *gin.Context) {
	p, perr := killp.FindProcessByName(c.Param("commit_version"))
	if perr != nil {
		fmt.Println(perr)
	}
	p.Kill()
}
*/