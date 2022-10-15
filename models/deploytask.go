package models

type DeployTask struct {
	ID                  int     `gorm:"primary_key"`
	ApplyId             int     `gorm:"type:int(11);not null;default:0"`
	GroupId             int     `gorm:"type:int(11);not null;default:0"`
	Status              int     `json:"status" gorm:"type:varchar(50);"`
	Content             string  `gorm:"type:text;not null"`
	CreateBy   			string `json:"create_by" gorm:"type:varchar(128);"` // 创建人
	UpdateBy   			string `json:"update_by" gorm:"type:varchar(128);"` // 更新者
}

func (m *DeployTask) TableName() string {
	return "dep_task"
}

func (m *DeployTask) List(query QueryParam) ([]DeployTask, bool) {
	var data []DeployTask
	ok := GetMulti(&data, query)
	return data, ok
}

func (m *DeployTask) GetByApplyId(id int) bool {
	return GetOne(m, QueryParam{
		Where: []WhereParam{
			WhereParam{
				Field: "apply_id",
				Prepare: id,
			},
		},
	})
}

func (m *DeployTask) UpdateByFields(data map[string]interface{}, query QueryParam) bool {
	return Update(m, data, query)
}

func (m *DeployTask) Create() bool {
	//m.Ctime = int(time.Now().Unix())
	return Create(m)
}

func (m *DeployTask) Delete(query QueryParam) bool {
	return Delete(m, query)
}
