package tools

import (
	"github.com/gin-gonic/gin"
	"vAdmin/models/tools"
	tools2 "vAdmin/tools"
	"vAdmin/tools/app"
	config2 "vAdmin/tools/config"
	"net/http"
)

func GetDBTableList(c *gin.Context) {
	var res app.Response
	var data tools.DBTables
	var err error
	var pageSize = 10
	var pageIndex = 1
	if config2.MysqlConfig.Dbtype=="sqlite3"{
		res.Msg="对不起，sqlite3 暂不支持代码生成！"
		c.JSON(http.StatusOK, res.ReturnError(500))
		return
	}

	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize = tools2.StrToInt(err, size)
	}

	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex = tools2.StrToInt(err, index)
	}

	data.TableName = c.Request.FormValue("tableName")
	result, count, err := data.GetPage(pageSize, pageIndex)
	tools2.HasError(err, "", -1)

	var mp = make(map[string]interface{}, 3)
	mp["list"] = result
	mp["count"] = count
	mp["pageIndex"] = pageIndex
	mp["pageSize"] = pageSize


	res.Data = mp

	c.JSON(http.StatusOK, res.ReturnOK())
}
