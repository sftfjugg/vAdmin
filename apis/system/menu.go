package system

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"vAdmin/models"
	"vAdmin/tools"
	"vAdmin/tools/app"
)

func GetMenuList(c *gin.Context) {
	var Menu models.Menu
	Menu.MenuName = c.Request.FormValue("menuName")
	Menu.Visible = c.Request.FormValue("visible")
	Menu.Title = c.Request.FormValue("title")
	Menu.DataScope = tools.GetUserIdStr(c)
	result, err := Menu.SetMenu()
	tools.HasError(err, "抱歉未找到相关信息", -1)

	app.OK(c, result, "")
}

func GetMenu(c *gin.Context) {
	var data models.Menu
	id, err := tools.StringToInt(c.Param("id"))
	data.MenuId = id
	result, err := data.GetByMenuId()
	tools.HasError(err, "抱歉未找到相关信息", -1)
	app.OK(c, result, "")
}

func GetMenuTreeRoleselect(c *gin.Context) {
	var Menu models.Menu
	var SysRole models.SysRole
	id, err := tools.StringToInt(c.Param("roleId"))
	SysRole.RoleId = id
	result, err := Menu.SetMenuLable()
	tools.HasError(err, "抱歉未找到相关信息", -1)
	menuIds := make([]int, 0)
	if id != 0 {
		menuIds, err = SysRole.GetRoleMeunId()
		tools.HasError(err, "抱歉未找到相关信息", -1)
	}
	app.Custum(c, gin.H{
		"code":        200,
		"menus":       result,
		"checkedKeys": menuIds,
	})
}

func GetMenuTreeelect(c *gin.Context) {
	var data models.Menu
	result, err := data.SetMenuLable()
	tools.HasError(err, "抱歉未找到相关信息", -1)
	app.OK(c, result, "")
}

func InsertMenu(c *gin.Context) {
	var data models.Menu
	err := c.BindWith(&data, binding.JSON)
	tools.HasError(err, "抱歉未找到相关信息", -1)
	data.CreateBy = tools.GetUserIdStr(c)
	result, err := data.Create()
	tools.HasError(err, "抱歉未找到相关信息", -1)
	app.OK(c, result, "")
}

func UpdateMenu(c *gin.Context) {
	var data models.Menu
	err2 := c.BindWith(&data, binding.JSON)
	data.UpdateBy = tools.GetUserIdStr(c)
	tools.HasError(err2, "修改失败", -1)
	_, err := data.Update(data.MenuId)
	tools.HasError(err, "", 501)
	app.OK(c, "", "修改成功")

}


func DeleteMenu(c *gin.Context) {
	var data models.Menu
	id, err := tools.StringToInt(c.Param("id"))
	data.UpdateBy = tools.GetUserIdStr(c)
	_, err = data.Delete(id)
	tools.HasError(err, "删除失败", 500)
	app.OK(c, "", "删除成功")
}

func GetMenuRole(c *gin.Context) {
	var Menu models.Menu
	result, err := Menu.SetMenuRole(tools.GetRoleName(c))
	tools.HasError(err, "获取失败", 500)
	app.OK(c, result, "")
}

func GetMenuIDS(c *gin.Context) {
	var data models.RoleMenu
	data.RoleName = c.GetString("role")
	data.UpdateBy = tools.GetUserIdStr(c)
	result, err := data.GetIDS()
	tools.HasError(err, "获取失败", 500)
	app.OK(c, result, "")
}
