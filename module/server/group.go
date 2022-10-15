// Copyright 2019 syncd Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package server

import (
    "errors"
    "fmt"
    "vAdmin/models"
    "vAdmin/tools/gois"
)

type Group struct {
    ID      int     `json:"id"`
    Name    string  `json:"name"`
}

func GroupGetMapByIds(ids []int) (map[int]Group, error) {
    if len(ids) == 0 {
        return nil, nil
    }
    group := &models.ServerGroup{}
    groupList, ok := group.List(models.QueryParam{
        Where: []models.WhereParam{
            models.WhereParam{
                Field: "id",
                Tag: "IN",
                Prepare: ids,
            },
        },
    })
    if !ok {
        return nil, errors.New("get server group list failed")
    }
    groupMap := make(map[int]Group)
    for _, l := range groupList{
        groupMap[l.ID] = Group{
            ID: l.ID,
            Name: l.Name,
        }
    }
    return groupMap, nil
}

func (g *Group) Create() error {
    serverGroup := models.ServerGroup{
        Name: g.Name,
    }
    if ok := serverGroup.Create(); !ok {
        return errors.New("create server group data failed")
    }
    return nil
}

func (g *Group) Update() error {
    serverGroup := models.ServerGroup{
        ID: g.ID,
        Name: g.Name,
    }
    if ok := serverGroup.Update(); !ok {
        return errors.New("update server group data failed")
    }
    return nil
}

func (g *Group) List(keyword string, offset, limit int) ([]Group, error) {
    group := models.ServerGroup{}
    list, ok := group.List(models.QueryParam{
        Fields: "id, name, ctime",
        Offset: offset,
        Limit: limit,
        Order: "id DESC",
        Where: g.parseWhereConds(keyword),
    })
    if !ok {
        return nil, errors.New("get server group list failed")
    }

    var groupList []Group
    for _, l := range list {
        groupList = append(groupList, Group{
            ID: l.ID,
            Name: l.Name,
        })
    }
    return groupList, nil
}

func (g *Group) Total(keyword string) (int, error) {
    group := models.ServerGroup{}
    total, ok := group.Count(models.QueryParam{
        Where: g.parseWhereConds(keyword),
    })
    if !ok {
        return 0, errors.New("get server group count failed")
    }
    return total, nil
}

func (g *Group) Delete() error {
    group := &models.ServerGroup{
        ID: g.ID,
    }
    if ok := group.Delete(); !ok {
        return errors.New("delete server group failed")
    }
    return nil
}

func (g *Group) Detail() error {
    group := models.ServerGroup{}
    if ok := group.Get(g.ID); !ok {
        return errors.New("get server group detail failed")
    }
    if group.ID == 0 {
        return errors.New("server group not exists")
    }

    g.ID = group.ID
    g.Name = group.Name

    return nil
}

func (g *Group) parseWhereConds(keyword string) []models.WhereParam {
    var where []models.WhereParam
    if keyword != "" {
        if gois.IsInteger(keyword) {
            where = append(where, models.WhereParam{
                Field: "id",
                Prepare: keyword,
            })
        } else {
            where = append(where, models.WhereParam{
                Field: "name",
                Tag: "LIKE",
                Prepare: fmt.Sprintf("%%%s%%", keyword),
            })
        }
    }
    return where
}
