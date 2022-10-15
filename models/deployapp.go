package models

import (
	"fmt"
	orm "vAdmin/database"
	//"vAdmin/pkg/logger"
)

type DeployApp struct {
	ID                  int     `json:"id" gorm:"type:int(11);unsigned;primary_key"`
	DemandId            string  `json:"demand_id" gorm:"type:varchar(100);not null;default:''"`
	CommitVersion       string  `json:"commit_version" gorm:"type:varchar(500);not null;default:''"`
	Description         string  `json:"description" gorm:"type:varchar(500);not null;default:''"`
	OnlineCluster       string  `json:"online_cluster" gorm:"type:varchar(500);not null;default:''"`
	DeployMode		    string	`json:"deploy_mode" gorm:"type:varchar(25);not null;default:'rolling'"`
	AuditStatus         string  `json:"audit_status" gorm:"type:varchar(25);not null;default:'rolling'"`
	AuditRefusalReasion string  `json:"audit_refusal_reasion" gorm:"type:varchar(500);not null;default:''"`
	AuditNotice			string  `json:"audit_notice" gorm:"type:varchar(50);"`
	DeployNotice		string  `json:"deploy_notice" gorm:"type:varchar(50);"`
	Status              string  `json:"status" gorm:"type:varchar(25);not null;default:'not_online'"`
	DeployUser          string  `json:"deploy_user" gorm:"type:varchar(100);not null;default:''"`
	CreateByname   		string `json:"create_byname" gorm:"type:varchar(25);"`
	CreateByemail  	    string `json:"create_byemail" gorm:"type:varchar(25);"`
	CreatePhone  	    string `json:"create_phone" gorm:"type:varchar(25);"`
	CreateBy   			string `json:"create_by" gorm:"type:varchar(128);"` // 创建人
	UpdateBy   			string `json:"update_by" gorm:"type:varchar(128);"` // 更新者
	// CreatedAt        string `json:"created_at" gorm:"type:timestamp;"` // grom自动生成
	// UpdatedAt	    string `json:"update_at" gorm:"type:timestamp;"` // grom自动生成
	// Daylabel  			string  `json:"daylabel" gorm:"type:varchar(10);"`

	DataScope string `json:"dataScope" gorm:"-"`
	Params    string `json:"params"  gorm:"-"`
	BaseModel
}

/*
type DeployAppManager struct{}

func NewDeployAppManager() *DeployAppManager {
	return &DeployAppManager{}
}

var DefaultDeployAppManager *DeployAppManager

func init() {
	DefaultDeployAppManager = NewDeployAppManager()
}

type Result struct {
	DemandID string
	CommitVersion string
}
var DeployInfo Result

func (h *DeployAppManager) GetDeployAppRecordByID(id int) (*Result, error) {
	//sql := `select demand_id, commit_version from dep_app where id = ?`
	//err := orm.Eloquent.Raw(&Result{}, sql, id).Scan(&DeployInfo)
	err := orm.Eloquent.Where("id = ?", id).First(&DeployInfo)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return &DeployInfo, nil
}
*/

type RollbackApp struct {
	ID            int     `json:"id" gorm:"type:int(11);unsigned;primary_key"`
	AppName       string  `json:"app_name" gorm:"type:varchar(25);not null;default:0"`
	RollVersion   string  `json:"roll_version" gorm:"type:varchar(100);not null;default:0"`
	DateVersion   string  `json:"date_version" gorm:"type:varchar(100);not null;default:0"`
	RollStatus    int     `json:"roll_status" gorm:"type:int(11);not null;default:0"`
	CreateBy   	  string `json:"create_by" gorm:"type:varchar(128);"` // 创建人
	UpdateBy   	  string `json:"update_by" gorm:"type:varchar(128);"` // 更新者

	DataScope string `json:"dataScope" gorm:"-"`
	Params    string `json:"params"  gorm:"-"`
	BaseModel
}

// 定义表名
func (DeployApp) TableName() string {
	return "dep_app"
}

// 定义表名
func (RollbackApp) RTableName() string {
	return "dep_rollback"
}

// 创建工单
func (e *DeployApp) Create() (DeployApp, error) {
	var doc DeployApp
	//e.updated_at = int(time.Now().Unix())
	result := orm.Eloquent.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 以满足条件的字段作为条件查询
func (e *DeployApp) Get() (DeployApp, error) {
	var doc DeployApp
	table := orm.Eloquent.Table(e.TableName())

	// 此if条件很关键，提供查看、修改、删除数据。
	if e.ID != 0  {
		table = table.Where("id = ?", e.ID)
	}

	if e.AuditStatus != ""  {
		table = table.Where("audit_status = ?", e.AuditStatus)
	}

	if e.Status != ""  {
		table = table.Where("status = ?", e.Status)
	}

	if e.DemandId != ""  {
		table = table.Where("demand_id = ?", e.DemandId)
	}

	if e.DeployMode != ""  {
		table = table.Where("deploy_mode = ?", e.DeployMode)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 以ID作为参数构造的查询
func (m *DeployApp) GetPram(id int) bool {
	return GetByPk(m, id)
}

func (m *DeployApp) List(query QueryParam) ([]DeployApp, bool) {
	var data []DeployApp
	ok := GetMulti(&data, query)
	return data, ok
}

func (m *DeployApp) UpdateByFields(data map[string]interface{}, query QueryParam) bool {
	return Update(m, data, query)
}


func (m *DeployApp) Count(query QueryParam) (int, bool) {
	var count int
	ok := Count(m, &count, query)
	return count, ok
}

// 条件查询获取RollbackApp
func (e *RollbackApp) GetRollVersion() (RollbackApp, error) {
	var doc RollbackApp
	table := orm.Eloquent.Table(e.RTableName())

	if e.AppName != ""  {
		table = table.Where("app_name = ?", e.AppName)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取已上线带分页
func (e *DeployApp) GetPage(pageSize int, pageIndex int) ([]DeployApp, int, error) {
	var doc []DeployApp

	fmt.Println(pageSize,pageIndex)
	table := orm.Eloquent.Select("*," +
		"case " +
		" when created_at is not null and DATEDIFF(NOW(),created_at)=0 then 'today'" +
		" when created_at is not null and DATEDIFF(NOW(),created_at)=1 then 'yesterday'" +
		" when created_at is not null and DATEDIFF(NOW(),created_at)=2 then 'byesterday'" +
		" when created_at is not null and DATEDIFF(NOW(),created_at)=8 then '7day'" +
		" when created_at is not null and DATEDIFF(NOW(),created_at)>0 and DATEDIFF(NOW(),created_at)<=30 then 'within_one_month'" +
		" when created_at is not null and DATEDIFF(NOW(),created_at)>0 and DATEDIFF(NOW(),created_at)<=90 then 'within_three_months'" +
		" when created_at is not null and DATEDIFF(NOW(),created_at)>0 and DATEDIFF(NOW(),created_at)<=365 then 'within_a_year'" +
		" else 'any_time' end as daylabel").Table(e.TableName())

	if e.AuditStatus != ""  {
		table = table.Where("audit_status = ?", e.AuditStatus)
	}

	if e.Status != ""  {
		table = table.Where("status = ?", e.Status)
	}

	if e.DemandId != ""  {
		table = table.Where("demand_id = ?", e.DemandId)
	}

	if e.DeployMode != ""  {
		table = table.Where("deploy_mode = ?", e.DeployMode)
	}

	var count int

	if err := table.Offset((pageIndex - 1) * pageSize).Order("DATE_FORMAT(created_at,'%Y-%m-%d') desc").Order("demand_id desc").Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	if err := table.Where("deleted_at is NULL").Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return doc, count, nil
}


func (e *RollbackApp) GetList() ([]RollbackApp, error) {
	var doc []RollbackApp

	table := orm.Eloquent.Select("date_version,roll_version").Table(e.RTableName())

	if e.DateVersion != ""  {
		table = table.Where("date_version = ?", e.DateVersion)
	}

	if e.AppName != ""  {
		table = table.Where("app_name = ?", e.AppName)
	}

	if err := table.Order("DATE_FORMAT(date_version,'%Y-%m-%d') desc").Limit(5).Find(&doc).Error; err != nil {
		return nil, err
	}

	return doc, nil
}

// 更新DeployApp
func (e *DeployApp) Update(id int) (update DeployApp, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id = ?", id).First(&update).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Table(e.TableName()).Model(&update).Updates(&e).Error; err != nil {
		return
	}
	return
}

func (m *DeployApp) Update2() bool {
	return UpdateByPk(m)
}