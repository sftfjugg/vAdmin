package system

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"vAdmin/models"
	"vAdmin/tools"
	"vAdmin/tools/app"
	"vAdmin/tools/app/msg"
	"net/http"
)

func GetConfigList(c *gin.Context) {
	var data models.SysConfig
	var err error
	var pageSize = 10
	var pageIndex = 1

	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize = tools.StrToInt(err, size)
	}

	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex = tools.StrToInt(err, index)
	}

	data.ConfigKey = c.Request.FormValue("configKey")
	data.ConfigName = c.Request.FormValue("configName")
	data.ConfigType = c.Request.FormValue("configType")
	data.DataScope = tools.GetUserIdStr(c)
	result, count, err := data.GetPage(pageSize, pageIndex)
	tools.HasError(err, "", -1)

	var mp = make(map[string]interface{}, 3)
	mp["list"] = result
	mp["count"] = count
	mp["pageIndex"] = pageIndex
	mp["pageSize"] = pageSize

	var res app.Response
	res.Data = mp

	c.JSON(http.StatusOK, res.ReturnOK())
}

func GetConfig(c *gin.Context) {
	var Config models.SysConfig
	Config.ConfigId, _ = tools.StringToInt(c.Param("configId"))
	result, err := Config.Get()
	tools.HasError(err, "抱歉未找到相关信息", -1)

	var res app.Response
	res.Data = result

	c.JSON(http.StatusOK, res.ReturnOK())
}


func GetConfigByConfigKey(c *gin.Context) {
	var Config models.SysConfig
	Config.ConfigKey = c.Param("configKey")
	result, err := Config.Get()
	tools.HasError(err, "抱歉未找到相关信息", -1)

	app.OK(c, result,result.ConfigValue)
}

func InsertConfig(c *gin.Context) {
	var data models.SysConfig
	err := c.BindWith(&data, binding.JSON)
	data.CreateBy = tools.GetUserIdStr(c)
	tools.HasError(err, "", 500)
	result, err := data.Create()
	tools.HasError(err, "", -1)

	app.OK(c, result,"")
}

func UpdateConfig(c *gin.Context) {
	var data models.SysConfig
	err := c.BindWith(&data, binding.JSON)
	tools.HasError(err, "数据解析失败", -1)
	data.UpdateBy = tools.GetUserIdStr(c)
	result, err := data.Update(data.ConfigId)
	tools.HasError(err, "", -1)
	app.OK(c, result,"")
}

func DeleteConfig(c *gin.Context) {
	var data models.SysConfig
	data.UpdateBy = tools.GetUserIdStr(c)
	IDS := tools.IdsStrToIdsIntGroup("configId", c)
	result, err := data.BatchDelete(IDS)
	tools.HasError(err, "修改失败", 500)
	app.OK(c, result, msg.DeletedSuccess)
}


