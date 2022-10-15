// Copyright 2019 syncd Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package deploy

import (
    "errors"
    "vAdmin/models"
)

//deploy_task
type Deploy struct{
    ID          int     `json:"id"`
    ApplyId     int     `json:"apply_id"`
    GroupId     int     `json:"group_id"`
    Status      int     `json:"status"`
    Content     string  `json:"content"`
    CreateBy    string  `json:"create_by"`
}

const (
    DEPLOY_STATUS_NONE = 0
    DEPLOY_STATUS_START = 1
    DEPLOY_STATUS_SUCCESS = 2
    DEPLOY_STATUS_FAILED = 3
)

func (d *Deploy) TaskList() ([]Deploy, error) {
    dt := &models.DeployTask{}
    list, ok := dt.List(models.QueryParam{
        Fields: "id, status, content, created_at",
        Where: []models.WhereParam{
            models.WhereParam{
                Field: "id",
                Prepare: d.ID,
            },
            models.WhereParam{
                Field: "group_id",
                Prepare: d.GroupId,
            },
        },
    })
    if !ok {
        return nil, errors.New("get deploy task list failed")
    }
    var (
        deployList []Deploy
    )
    for _, l := range list {
        deployList = append(deployList, Deploy{
            ID: l.ID,
            ApplyId: l.ApplyId,
            GroupId: l.GroupId,
            Status: l.Status,
            Content: l.Content,
            CreateBy: l.CreateBy,
        })
    }
    return deployList, nil
}

func (d *Deploy) Create() error {
    deploy := models.DeployTask{
        ApplyId: d.ApplyId,
        GroupId: d.GroupId,
        Status: d.Status,
    }
    if ok := deploy.Create(); !ok {
        return errors.New("create deploy task failed")
    }
    return nil
}

func (d *Deploy) UpdateStatus() error {
    dt := &models.DeployTask{}
    updateData := map[string]interface{}{
        "status": d.Status,
    }
    if ok := dt.UpdateByFields(updateData,models.QueryParam{
        Where: []models.WhereParam{
            models.WhereParam{
                Field: "apply_id",
                Prepare: d.ApplyId,
            },
            models.WhereParam{
                Field: "group_id",
                Prepare: d.GroupId,
            },
        },
    }); !ok {
        return errors.New("update deploy task result failed")
    }
    return nil
}

func (d *Deploy) UpdateResult() error {
    dt := &models.DeployTask{}
    updateData := map[string]interface{}{
        "status": d.Status,
        "content": d.Content,
    }
    if ok := dt.UpdateByFields(updateData, models.QueryParam{
        Where: []models.WhereParam{
            models.WhereParam{
                Field: "apply_id",
                Prepare: d.ApplyId,
            },
            models.WhereParam{
                Field: "group_id",
                Prepare: d.GroupId,
            },
        },
    }); !ok {
        return errors.New("update deploy task result failed")
    }
    return nil
}

func (d *Deploy) DeleteByApplyId() error {
    dep := &models.DeployTask{
        ApplyId: d.ApplyId,
    }
    if ok := dep.Delete(models.QueryParam{
        Where: []models.WhereParam{
            models.WhereParam{
                Field: "apply_id",
                Prepare: d.ApplyId,
            },
        },
    }); !ok {
        return errors.New("remove deploy task failed")
    }
    return nil
}