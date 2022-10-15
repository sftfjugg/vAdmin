// Copyright 2019 syncd Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

type Server struct {
    ID          int     `gorm:"primary_key"`
    GroupId     int     `gorm:"type:int(11);not null;default:0"`
    Name        string  `gorm:"type:varchar(100);not null;default:''"`
    Ip          string  `gorm:"type:varchar(100);not null;default:''"`
    SSHPort     int     `gorm:"type:int(11);not null;default:0"`
    CreateBy   			string `json:"create_by" gorm:"type:varchar(128);"` // 创建人
    UpdateBy   			string `json:"update_by" gorm:"type:varchar(128);"` // 更新者
}

func (m *Server) TableName() string {
    return "dep_server"
}

func (m *Server) Create() bool {
    //m.Ctime = int(time.Now().Unix())
    return Create(m)
}

func (m *Server) Update() bool {
    return UpdateByPk(m)
}

func (m *Server) List(query QueryParam) ([]Server, bool) {
    var data []Server
    ok := GetMulti(&data, query)
    return data, ok
}

func (m *Server) Count(query QueryParam) (int, bool) {
    var count int
    ok := Count(m, &count, query)
    return count, ok
}

func (m *Server) Delete() bool {
    return DeleteByPk(m)
}

func (m *Server) Get(id int) bool {
    return GetByPk(m, id)
}
