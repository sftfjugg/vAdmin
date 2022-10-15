package gorm

import (
	"github.com/jinzhu/gorm"
	"vAdmin/models"
	"vAdmin/models/process"
	"vAdmin/models/tools"
)

func AutoMigrate(db *gorm.DB) error {
	db.SingularTable(true)
	return db.AutoMigrate(

		// 系统管理
		new(models.CasbinRule),
		new(tools.SysTables),
		new(tools.SysColumns),
		new(models.Dept),
		new(models.Menu),
		new(models.LoginLog),
		new(models.SysOperLog),
		new(models.RoleMenu),
		new(models.SysRoleDept),
		new(models.SysUser),
		new(models.SysRole),
		new(models.Post),
		new(models.DictData),
		new(models.SysConfig),
		new(models.DictType),

		// 流程中心
		new(process.Classify),
		new(process.TplInfo),
		new(process.TplData),
		new(process.WorkOrderInfo),
		new(process.TaskInfo),
		new(process.Info),
		new(process.History),
		new(process.CirculationHistory),
	).Error
}
