package dict

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"vAdmin/models"
	"vAdmin/tools"
	"vAdmin/tools/app"
	"net/http"
)

func GetDictTypeList(c *gin.Context) {
	var data models.DictType
	var err error
	var pageSize = 10
	var pageIndex = 1

	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize = tools.StrToInt(err, size)
	}

	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex = tools.StrToInt(err, index)
	}

	data.DictName = c.Request.FormValue("dictName")
	id := c.Request.FormValue("dictId")
	data.DictId, _ = tools.StringToInt(id)
	data.DictType = c.Request.FormValue("dictType")
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

func GetDictType(c *gin.Context) {
	var DictType models.DictType
	DictType.DictName = c.Request.FormValue("dictName")
	DictType.DictId, _ = tools.StringToInt(c.Param("dictId"))
	result, err := DictType.Get()
	tools.HasError(err, "抱歉未找到相关信息", -1)
	var res app.Response
	res.Data = result
	c.JSON(http.StatusOK, res.ReturnOK())
}

func GetDictTypeOptionSelect(c *gin.Context) {
	var DictType models.DictType
	DictType.DictName = c.Request.FormValue("dictName")
	DictType.DictId, _ = tools.StringToInt(c.Param("dictId"))
	result, err := DictType.GetList()
	tools.HasError(err, "抱歉未找到相关信息", -1)
	var res app.Response
	res.Data = result
	c.JSON(http.StatusOK, res.ReturnOK())
}

func InsertDictType(c *gin.Context) {
	var data models.DictType
	err := c.BindWith(&data, binding.JSON)
	data.CreateBy = tools.GetUserIdStr(c)
	tools.HasError(err, "", 500)
	result, err := data.Create()
	tools.HasError(err, "", -1)
	var res app.Response
	res.Data = result
	c.JSON(http.StatusOK, res.ReturnOK())
}

func UpdateDictType(c *gin.Context) {
	var data models.DictType
	err := c.BindWith(&data, binding.JSON)
	data.UpdateBy = tools.GetUserIdStr(c)
	tools.HasError(err, "", -1)
	result, err := data.Update(data.DictId)
	tools.HasError(err, "", -1)
	var res app.Response
	res.Data = result
	c.JSON(http.StatusOK, res.ReturnOK())
}

func DeleteDictType(c *gin.Context) {
	var data models.DictType
	data.UpdateBy = tools.GetUserIdStr(c)
	IDS := tools.IdsStrToIdsIntGroup("dictId", c)
	result, err := data.BatchDelete(IDS)
	tools.HasError(err, "修改失败", 500)
	app.OK(c, result, "删除成功")
}
