package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"vAdmin/models"
	"vAdmin/tools/app"
	"net/http"
)

func GetRoleMenu(c *gin.Context) {
	var Rm models.RoleMenu
	err := c.ShouldBind(&Rm)
	result, err := Rm.Get()
	var res app.Response
	if err != nil {
		res.Msg = "抱歉未找到相关信息"
		c.JSON(http.StatusOK, res.ReturnError(200))
		return
	}
	res.Data = result
	c.JSON(http.StatusOK, res.ReturnOK())
}

type RoleMenuPost struct {
	RoleId   string
	RoleMenu []models.RoleMenu
}

func InsertRoleMenu(c *gin.Context) {

	var res app.Response
	res.Msg = "添加成功"
	c.JSON(http.StatusOK, res.ReturnOK())
	return

}

func DeleteRoleMenu(c *gin.Context) {
	var t models.RoleMenu
	id := c.Param("id")
	menuId := c.Request.FormValue("menu_id")
	fmt.Println(menuId)
	_, err := t.Delete(id, menuId)
	if err != nil {
		var res app.Response
		res.Msg = "删除失败"
		c.JSON(http.StatusOK, res.ReturnError(200))
		return
	}
	var res app.Response
	res.Msg = "删除成功"
	c.JSON(http.StatusOK, res.ReturnOK())
	return
}
