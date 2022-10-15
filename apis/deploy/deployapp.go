package deploy

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"sync"
	"vAdmin/models"
	"vAdmin/tools"
	"vAdmin/tools/app"
)

const (
	STATUS_INIT = 0
	STATUS_ING = 1
	STATUS_DONE = 2
	STATUS_FAILED = 3
)

const (
	DEPLOY_PARALLEL = "parallel"
	DEPLOY_SERIAL = "rolling"
)

type Deploy struct {
	ID              int
	User            string
	Cmd          	string
	srvs            []*Server
	status          int
	wg              sync.WaitGroup
}

type DeployResult struct {
	ID          int
	Status      int
	ServerRest  []*ServerResult
}

func GetDeployAppList(c *gin.Context) {
	var data models.DeployApp
	var err error
	var pageSize = 10
	var pageIndex = 1
	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize = tools.StrToInt(err, size)
	}
	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex = tools.StrToInt(err, index)
	}

	data.DemandId = c.Request.FormValue("demand_id")
	data.CommitVersion = c.Request.FormValue("commit_version")
	data.DeployMode = c.Request.FormValue("deploy_mode")
	data.Description = c.Request.FormValue("description")
	data.AuditStatus = c.Request.FormValue("audit_status")
	data.AuditRefusalReasion = c.Request.FormValue("audit_refusal_reasion")
	data.Status = c.Request.FormValue("status")
	data.CreateByname = c.Request.FormValue("create_byname")
	data.CreateByemail = c.Request.FormValue("create_byemail")
	data.CreatePhone = c.Request.FormValue("create_phone")

	// data.Daylabel = c.Request.FormValue("daylabel")
	// data.CreatedAt = c.Request.FormValue("created_at") // grom自动生成
	data.DataScope = tools.GetUserIdStr(c)
	result, count, err := data.GetPage(pageSize, pageIndex)
	tools.HasError(err, "", -1)

	app.PageOK(c, result, count, pageIndex, pageSize, "")
}

func GetRollbackAppList(c *gin.Context) {
	var data models.RollbackApp
	var err error

	data.AppName = c.Request.FormValue("app_name")
	data.RollVersion = c.Request.FormValue("roll_version")
	data.DateVersion = c.Request.FormValue("date_version")
	data.RollStatus = tools.StrToInt(err, c.Request.FormValue("roll_status"))

	data.DataScope = tools.GetUserIdStr(c)
	result, err := data.GetList()
	tools.HasError(err, "", -1)

	app.OK(c, result, "")
}

func GetDeployApp(c *gin.Context) {
	var data models.DeployApp
	data.ID, _ = tools.StringToInt(c.Param("id"))
	result, err := data.Get()
	tools.HasError(err, "抱歉未找到相关信息", -1)

	app.OK(c, result, "")
}

func InsertDeployApp(c *gin.Context) {
	var data models.DeployApp
	err := c.ShouldBindJSON(&data)
	data.CreateBy = tools.GetUserIdStr(c)
	tools.HasError(err, "", 500)
	result, err := data.Create()
	tools.HasError(err, "", -1)
	app.OK(c, result, "")
}

func UpdateDeployApp(c *gin.Context) {
	var data models.DeployApp
	err := c.BindWith(&data, binding.JSON)
	tools.HasError(err, "数据解析失败", -1)
	data.UpdateBy = tools.GetUserIdStr(c)
	result, err := data.Update(data.ID)
	tools.HasError(err, "", -1)
	app.OK(c, result, "")
}

func (d *Deploy) Parallel() {
	if d.status == STATUS_FAILED {
		return
	}
	d.status = STATUS_ING
	status := STATUS_DONE
	for _, srv := range d.srvs {
		d.wg.Add(1)
		go func() {
			if d.status == STATUS_ING {
				srv.Deploy()
				if srv.Result().Status == STATUS_FAILED {
					status = STATUS_FAILED
				}
			}
			defer d.wg.Done()
		}()
	}
	d.wg.Wait()
	d.status = status
}

func (d *Deploy) Serial() {
	if d.status == STATUS_FAILED {
		return
	}
	d.status = STATUS_ING
	status := STATUS_DONE
	for _, srv := range d.srvs {
		if d.status == STATUS_ING {
			srv.Deploy()
			if srv.Result().Status == STATUS_FAILED {
				status = STATUS_FAILED
			}
		}
	}
	d.status = status
}

func (d *Deploy) Result() ([]*ServerResult, int) {
	var rest []*ServerResult
	for _, srv := range d.srvs {
		rest = append(rest, srv.Result())
	}
	return rest, d.status
}


func (d *Deploy) Terminate() {
	d.status = STATUS_FAILED
	for _, srv := range d.srvs {
		srv.Terminate()
	}
}

func (d *Deploy) AddServer(id int, addr string, port int) {
	srv := &Server{
		ID: id,
		Addr: addr,
		User: d.User,
		Port: port,
		Cmd: d.Cmd,
	}
	NewServer(srv)
	d.srvs = append(d.srvs, srv)
}