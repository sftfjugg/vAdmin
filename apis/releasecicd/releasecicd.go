package releasecicd

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"vAdmin/apis/common"
	"vAdmin/models"
	"vAdmin/tools"
	"vAdmin/tools/app"
	"vAdmin/tools/app/msg"
)

// @Summary 获取已发生产的需求
// @Description 获取JSON
// @Tags 已发生产
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/releasecicdList [get]
// @Security Bearer
func GetReleasecicdList(c *gin.Context) {
	var data models.Releasecicd
	var err error
	var pageSize = 10
	var pageIndex = 1
	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize = tools.StrToInt(err, size)
	}
	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex = tools.StrToInt(err, index)
	}

	data.Productname = c.Request.FormValue("productname")
	data.Rid = c.Request.FormValue("rid")
	data.Fid = c.Request.FormValue("fid")
	data.Bid = c.Request.FormValue("bid")
	data.Appname = c.Request.FormValue("appname")
	data.Commitid = c.Request.FormValue("commitid")
	data.Prdversion = c.Request.FormValue("prdversion")
	data.Openedby = c.Request.FormValue("openedby")
	data.Openeddate = c.Request.FormValue("openeddate")
	data.Cicdtime = c.Request.FormValue("cicdtime")
	data.Deploytime = c.Request.FormValue("deploytime")
	data.Fidstatus = c.Request.FormValue("fidstatus")

	data.Daylabel = c.Request.FormValue("daylabel")
	data.DataScope = tools.GetUserIdStr(c)
	result, count, err := data.GetPage(pageSize, pageIndex)
	tools.HasError(err, "", -1)

	app.PageOK(c, result, count, pageIndex, pageSize, "")
}

func GetReleasecicd(c *gin.Context) {
	var data models.Releasecicd
	data.Id, _ = tools.StringToInt(c.Param("id"))
	result, err := data.Get()
	tools.HasError(err, "抱歉未找到相关信息", -1)

	app.OK(c, result, "")
}

func InsertReleasecicd(c *gin.Context) {
	var data models.Releasecicd
	err := c.ShouldBindJSON(&data)
	data.CreateBy = tools.GetUserIdStr(c)
	tools.HasError(err, "", 500)
	result, err := data.Create()
	tools.HasError(err, "", -1)
	app.OK(c, result, "")
}

func UpdateReleasecicd(c *gin.Context) {
	var data models.Releasecicd
	err := c.BindWith(&data, binding.JSON)
	tools.HasError(err, "数据解析失败", -1)
	data.UpdateBy = tools.GetUserIdStr(c)
	result, err := data.Update(data.Id)
	tools.HasError(err, "", -1)

	app.OK(c, result, "")
}

func DeleteReleasecicd(c *gin.Context) {
	var data models.Releasecicd
	data.UpdateBy = tools.GetUserIdStr(c)

	IDS := tools.IdsStrToIdsIntGroup("id", c)
	_, err := data.BatchDelete(IDS)
	tools.HasError(err, msg.DeletedFail, 500)
	app.OK(c, nil, msg.DeletedSuccess)
}

func GetRidList(c *gin.Context) {
	var data models.Releaserid
	var err error
	var pageSize = 10
	var pageIndex = 1
	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize = tools.StrToInt(err, size)
	}
	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex = tools.StrToInt(err, index)
	}

	data.Rid = c.Request.FormValue("rid")
	data.Rnotesadd = c.Request.FormValue("rnotesadd")
	data.Openeddate = c.Request.FormValue("openeddate")

	data.DataScope = tools.GetUserIdStr(c)
	result, count, err := data.GetPageAdd(pageSize, pageIndex)
	tools.HasError(err, "", -1)

	app.PageOK(c, result, count, pageIndex, pageSize, "")
}

// @Summary 获取未提测的需求
// @Description 获取JSON
// @Tags 未提测
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/releasetodoList [get]
// @Security Bearer
func GetReleasetodolist(c *gin.Context) {
	var data models.Releasecicd
	var err error
	var pageSize = 10
	var pageIndex = 1
	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize = tools.StrToInt(err, size)
	}
	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex = tools.StrToInt(err, index)
	}

	data.Rid = c.Request.FormValue("rid")
	data.Fid = c.Request.FormValue("fid")
	data.Bid = c.Request.FormValue("bid")
	data.Openedby = c.Request.FormValue("openedby")
	data.Openeddate = c.Request.FormValue("openeddate")
	data.Fidstatus = c.Request.FormValue("fidstatus")
	data.Sure_status = c.Request.FormValue("sure_status")
	data.Sure_time = c.Request.FormValue("sure_time")
	data.Sure_reasion = c.Request.FormValue("sure_reasion")

	data.Daylabel = c.Request.FormValue("daylabel")
	data.DataScope = tools.GetUserIdStr(c)
	result, count, err := data.GetPageTodo(pageSize, pageIndex)
	tools.HasError(err, "", -1)

	app.PageOK(c, result, count, pageIndex, pageSize, "")
}

// @Summary 获取已提测的需求
// @Description 获取JSON
// @Tags 已提测
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/releasestgList [get]
// @Security Bearer
func GetReleaseStglist(c *gin.Context) {
	var data models.Releasecicd
	var err error
	var pageSize = 10
	var pageIndex = 1
	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize = tools.StrToInt(err, size)
	}
	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex = tools.StrToInt(err, index)
	}

	data.Rid = c.Request.FormValue("rid")
	data.Fid = c.Request.FormValue("fid")
	data.Bid = c.Request.FormValue("bid")
	data.Openedby = c.Request.FormValue("openedby")
	data.Openeddate = c.Request.FormValue("openeddate")
	data.Cicdtime = c.Request.FormValue("cicdtime")
	data.Fidstatus = c.Request.FormValue("fidstatus")

	data.Daylabel = c.Request.FormValue("daylabel")
	data.DataScope = tools.GetUserIdStr(c)
	result, count, err := data.GetPageStg(pageSize, pageIndex)
	tools.HasError(err, "", -1)

	app.PageOK(c, result, count, pageIndex, pageSize, "")
}

func UpdateReleasetodo(c *gin.Context) {
	var data models.Releasecicd
	err := c.BindWith(&data, binding.JSON)
	tools.HasError(err, "数据解析失败", -1)
	data.UpdateBy = tools.GetUserIdStr(c)
	result, err := data.Update(data.Id)
	tools.HasError(err, "", -1)

	app.OK(c, result, "")
}

// @Summary 获取每天上线发布数趋势
// @Description 获取JSON
// @Tags 发布统计
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/releasechart/deploybyday [get]
func GetDeployByDay(c *gin.Context) {
	var data models.Releasecicd
	var err error

	data.DataScope = tools.GetUserIdStr(c)
	result, err := data.DeployByDay()
	tools.HasError(err, "", -1)

	app.OK(c,result,msg.GetSuccess)
}

// @Summary 获取每小时上线发布数趋势
// @Description 获取JSON
// @Tags 发布统计
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/releasechart/deployrate [get]
func GetDeployByHour(c *gin.Context) {
	var data models.Releasecicd
	var err error

	data.DataScope = tools.GetUserIdStr(c)
	result, err := data.DeploybyHour()
	tools.HasError(err, "", -1)

	app.OK(c,result,msg.GetSuccess)
}

// 分页数据
func  GetPageList(c *gin.Context) {
	page := common.GetPageIndex(c)
	limit := common.GetPageLimit(c)
	sort := common.GetPageSort(c)
	key := common.GetPageKey(c)
	//status := common.GetQueryToUint(c, "status")
	fid := common.GetQueryToStr(c,"fid")
	var whereOrder []models.PageWhereOrder
	order := "ID DESC"
	if len(sort) >= 2 {
		orderType := sort[0:1]
		order = sort[1:len(sort)]
		if orderType == "+" {
			order += " ASC"
		} else {
			order += " DESC"
		}
	}
	whereOrder = append(whereOrder, models.PageWhereOrder{Order: order})
	if key != "" {
		v := "%" + key + "%"
		var arr []interface{}
		arr = append(arr, v)
		arr = append(arr, v)
		whereOrder = append(whereOrder, models.PageWhereOrder{Where: "user_name like ? or real_name like ?", Value: arr})
	}
	if fid != "" {
		var arr []interface{}
		arr = append(arr, fid)
		whereOrder = append(whereOrder, models.PageWhereOrder{Where: "fid = ?", Value: arr})
	}
	var total uint64
	list:= []models.Releasecicd{}
	err := models.GetPageCom(&models.Releasecicd{}, &models.Releasecicd{}, &list, page, limit, &total, whereOrder...)
	if err != nil {
		common.ResErrSrv(c, err)
		return
	}
	common.ResSuccessPage(c, total, &list)
}