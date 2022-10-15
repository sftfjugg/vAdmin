package models

type ServerGroup struct {
	ID      int         `gorm:"primary_key"`
	Name    string      `gorm:"type:varchar(100);not null;default:''"`
	CreateBy   			string `json:"create_by" gorm:"type:varchar(128);"` // 创建人
	UpdateBy   			string `json:"update_by" gorm:"type:varchar(128);"` // 更新者
}

func (m *ServerGroup) TableName() string {
	return "dep_server_group"
}

func (m *ServerGroup) Create() bool {
	//m.Ctime = int(time.Now().Unix())
	return Create(m)
}

func (m *ServerGroup) Update() bool {
	return UpdateByPk(m)
}

func (m *ServerGroup) List(query QueryParam) ([]ServerGroup, bool) {
	var data []ServerGroup
	ok := GetMulti(&data, query)
	return data, ok
}

func (m *ServerGroup) Count(query QueryParam) (int, bool) {
	var count int
	ok := Count(m, &count, query)
	return count, ok
}

func (m *ServerGroup) Delete() bool {
	return DeleteByPk(m)
}

func (m *ServerGroup) Get(id int) bool {
	return GetByPk(m, id)
}
