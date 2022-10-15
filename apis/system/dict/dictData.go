package dict

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"vAdmin/models"
	"vAdmin/tools"
	"vAdmin/tools/app"
	"net/http"
)

func GetDictDataList(c *gin.Context) {
	var data models.DictData
	var err error
	var pageSize = 10
	var pageIndex = 1

	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize = tools.StrToInt(err, size)
	}

	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex = tools.StrToInt(err, index)
	}

	data.DictLabel = c.Request.FormValue("dictLabel")
	data.Status = c.Request.FormValue("status")
	data.DictType = c.Request.FormValue("dictType")
	id := c.Request.FormValue("dictCode")
	data.DictCode, _ = tools.StringToInt(id)
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

func GetDictData(c *gin.Context) {
	var DictData models.DictData
	DictData.DictLabel = c.Request.FormValue("dictLabel")
	DictData.DictCode, _ = tools.StringToInt(c.Param("dictCode"))
	result, err := DictData.GetByCode()
	tools.HasError(err, "抱歉未找到相关信息", -1)
	var res app.Response
	res.Data = result
	c.JSON(http.StatusOK, res.ReturnOK())
}

func GetDictDataByDictType(c *gin.Context) {
	var DictData models.DictData
	DictData.DictType = c.Param("dictType")
	result, err := DictData.Get()
	tools.HasError(err, "抱歉未找到相关信息", -1)

	var res app.Response
	res.Data = result
	c.JSON(http.StatusOK, res.ReturnOK())
}

func InsertDictData(c *gin.Context) {
	var data models.DictData
	err := c.BindWith(&data, binding.JSON)
	data.CreateBy = tools.GetUserIdStr(c)
	tools.HasError(err, "", 500)
	result, err := data.Create()
	tools.HasError(err, "", -1)
	var res app.Response
	res.Data = result
	c.JSON(http.StatusOK, res.ReturnOK())
}

func UpdateDictData(c *gin.Context) {
	var data models.DictData
	err := c.BindWith(&data, binding.JSON)
	data.UpdateBy = tools.GetUserIdStr(c)
	tools.HasError(err, "", -1)
	result, err := data.Update(data.DictCode)
	tools.HasError(err, "", -1)
	var res app.Response
	res.Data = result
	c.JSON(http.StatusOK, res.ReturnOK())
}

func DeleteDictData(c *gin.Context) {
	var data models.DictData
	data.UpdateBy = tools.GetUserIdStr(c)
	IDS := tools.IdsStrToIdsIntGroup("dictCode", c)
	result, err := data.BatchDelete(IDS)
	tools.HasError(err, "修改失败", 500)
	app.OK(c,result,"删除成功")
}
