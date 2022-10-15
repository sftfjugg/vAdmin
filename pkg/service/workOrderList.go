package service

import (
	"encoding/json"
	"fmt"
	//orm "vAdmin/database"
	orm "vAdmin/database"
	system "vAdmin/models"
	"vAdmin/models/process"
	"vAdmin/pkg/pagination"
	"vAdmin/tools"

	"github.com/gin-gonic/gin"
)

/*
  @Author : Rongxin Linghu
  @todo: 添加新的处理人时候，需要修改（先完善功能，后续有时间的时候优化一下这部分。）
*/

type WorkOrder struct {
	Classify int
	GinObj   *gin.Context
}

type workOrderInfo struct {
	process.WorkOrderInfo
	Principals   string `json:"principals"`
	StateName    string `json:"state_name"`
	DataClassify int    `json:"data_classify"`
}

func (w *WorkOrder) PureWorkOrderList() (result interface{}, err error) {
	var (
		workOrderInfoList []workOrderInfo
		processorInfo     system.SysUser
	)

	personSelectValue := "(JSON_CONTAINS(state, JSON_OBJECT('processor', %v)) and JSON_CONTAINS(state, JSON_OBJECT('process_method', 'person')))"
	roleSelectValue := "(JSON_CONTAINS(state, JSON_OBJECT('processor', %v)) and JSON_CONTAINS(state, JSON_OBJECT('process_method', 'role')))"
	departmentSelectValue := "(JSON_CONTAINS(state, JSON_OBJECT('processor', %v)) and JSON_CONTAINS(state, JSON_OBJECT('process_method', 'department')))"

	title := w.GinObj.DefaultQuery("title", "")
	startTime := w.GinObj.DefaultQuery("startTime", "")
	endTime := w.GinObj.DefaultQuery("endTime", "")
	isEnd := w.GinObj.DefaultQuery("isEnd", "")
	processor := w.GinObj.DefaultQuery("processor", "")
	priority := w.GinObj.DefaultQuery("priority", "")
	db := orm.Eloquent.Model(&process.WorkOrderInfo{}).
		Where("title like ?", fmt.Sprintf("%%%v%%", title))
	if startTime != "" {
		db = db.Where("create_time >= ?", startTime)
	}
	if endTime != "" {
		db = db.Where("create_time <= ?", endTime)
	}
	if isEnd != "" {
		db = db.Where("is_end = ?", isEnd)
	}
	if processor != "" && w.Classify != 1 {
		err = orm.Eloquent.Model(&processorInfo).
			Where("user_id = ?", processor).
			Find(&processorInfo).Error
		if err != nil {
			return
		}
		db = db.Where(fmt.Sprintf("(%v or %v or %v) and is_end = 0",
			fmt.Sprintf(personSelectValue, processorInfo.UserId),
			fmt.Sprintf(roleSelectValue, processorInfo.RoleId),
			fmt.Sprintf(departmentSelectValue, processorInfo.DeptId),
		))
	}
	if priority != "" {
		db = db.Where("priority = ?", priority)
	}

	// 获取当前用户信息
	switch w.Classify {
	case 1:
		// 待办工单
		// 1. 个人
		personSelect := fmt.Sprintf(personSelectValue, tools.GetUserId(w.GinObj))

		// 2. 角色
		roleSelect := fmt.Sprintf(roleSelectValue, tools.GetRoleId(w.GinObj))

		// 3. 部门
		var userInfo system.SysUser
		err = orm.Eloquent.Model(&system.SysUser{}).
			Where("user_id = ?", tools.GetUserId(w.GinObj)).
			Find(&userInfo).Error
		if err != nil {
			return
		}
		departmentSelect := fmt.Sprintf(departmentSelectValue, userInfo.DeptId)

		// 4. 变量会转成个人数据
		//db = db.Where(fmt.Sprintf("(%v or %v or %v or %v) and is_end = 0", personSelect, personGroupSelect, departmentSelect, variableSelect))
		db = db.Where(fmt.Sprintf("(%v or %v or %v) and is_end = 0", personSelect, roleSelect, departmentSelect))
	case 2:
		// 我创建的
		db = db.Where("creator = ?", tools.GetUserId(w.GinObj))
	case 3:
		// 我相关的
		db = db.Where(fmt.Sprintf("JSON_CONTAINS(related_person, '%v')", tools.GetUserId(w.GinObj)))
	case 4:
	// 所有工单
	default:
		return nil, fmt.Errorf("请确认查询的数据类型是否正确")
	}

	result, err = pagination.Paging(&pagination.Param{
		C:  w.GinObj,
		DB: db,
	}, &workOrderInfoList)
	if err != nil {
		err = fmt.Errorf("查询工单列表失败，%v", err.Error())
		return
	}
	return
}

func (w *WorkOrder) WorkOrderList() (result interface{}, err error) {

	var (
		principals        string
		StateList         []map[string]interface{}
		workOrderInfoList []workOrderInfo
		minusTotal        int
	)

	result, err = w.PureWorkOrderList()
	if err != nil {
		return
	}

	for i, v := range *result.(*pagination.Paginator).Data.(*[]workOrderInfo) {
		var (
			stateName    string
			structResult map[string]interface{}
			authStatus   bool
		)
		err = json.Unmarshal(v.State, &StateList)
		if err != nil {
			err = fmt.Errorf("json反序列化失败，%v", err.Error())
			return
		}
		if len(StateList) != 0 {
			// 仅待办工单需要验证
			// todo：还需要找最优解决方案
			if w.Classify == 1 {
				structResult, err = ProcessStructure(w.GinObj, v.Process, v.Id)
				if err != nil {
					return
				}

				authStatus, err = JudgeUserAuthority(w.GinObj, v.Id, structResult["workOrder"].(WorkOrderData).CurrentState)
				if err != nil {
					return
				}
				if !authStatus {
					minusTotal += 1
					continue
				}
			} else {
				authStatus = true
			}

			processorList := make([]int, 0)
			if len(StateList) > 1 {
				for _, s := range StateList {
					for _, p := range s["processor"].([]interface{}) {
						if int(p.(float64)) == tools.GetUserId(w.GinObj) {
							processorList = append(processorList, int(p.(float64)))
						}
					}
					if len(processorList) > 0 {
						stateName = s["label"].(string)
						break
					}
				}
			}
			if len(processorList) == 0 {
				for _, v := range StateList[0]["processor"].([]interface{}) {
					processorList = append(processorList, int(v.(float64)))
				}
				stateName = StateList[0]["label"].(string)
			}
			principals, err = GetPrincipal(processorList, StateList[0]["process_method"].(string))
			if err != nil {
				err = fmt.Errorf("查询处理人名称失败，%v", err.Error())
				return
			}
		}
		workOrderDetails := *result.(*pagination.Paginator).Data.(*[]workOrderInfo)
		workOrderDetails[i].Principals = principals
		workOrderDetails[i].StateName = stateName
		workOrderDetails[i].DataClassify = v.Classify
		if authStatus {
			workOrderInfoList = append(workOrderInfoList, workOrderDetails[i])
		}
	}

	result.(*pagination.Paginator).Data = &workOrderInfoList
	result.(*pagination.Paginator).TotalCount -= minusTotal

	return result, nil
}
