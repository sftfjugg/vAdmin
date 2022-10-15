// Copyright 2019 syncd Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package deploy

import (
	"errors"
	"time"
	"vAdmin/models"
)

type Apply struct {
    ID                  int     `json:"id"`
    DemandId            string  `json:"demand_id"`
    OnlineCluster       []int   `json:"online_cluster"`
    CommitVersion       string  `json:"commit_version"`
    Description         string  `json:"description"`
    AuditStatus         string  `json:"audit_status"`
    AuditRefusalReasion string  `json:"audit_refusal_reasion"`
    Status              string  `json:"status"`
    DeployUser          string  `json:"deploy_user"`
    DeployMode          string  `json:"deploy_mode"`
    CreateBy            string  `json:"create_by"`
    UpdateBy            string  `json:"update_by"`
    CreateByname        string  `json:"create_byname"`
    CreatePhone         string  `json:"create_phone"`
    //RollbackStatus      string  `json:"rollback_status"`
}

const (
    AUDIT_STATUS_PENDING = "unaudit"
    AUDIT_STATUS_OK = "audit_pass"
    AUDIT_STATUS_REFUSE = "audit_denied"
)

const (
    APPLY_STATUS_NONE = "not_online"
    APPLY_STATUS_ING = "onlineing"
    APPLY_STATUS_SUCCESS = "online_success"
    APPLY_STATUS_FAILED = "online_failed"
    APPLY_STATUS_DROP = "deprecated"
    APPLY_STATUS_ROLLBACK = "rollback"
)

func (a *Apply) Detail() error {
    apply := &models.DeployApp{}
    if ok := apply.GetPram(a.ID); !ok {
        return errors.New("get deploy apply detail failed")
    }
    if apply.ID == 0 {
        return errors.New("deploy apply detail not exists")
    }
    a.DemandId = apply.DemandId
    a.Description = apply.Description
    a.CommitVersion = apply.CommitVersion
    a.AuditStatus = apply.AuditStatus
    a.Status = apply.Status
    a.DeployMode = apply.DeployMode
    a.CreateBy = apply.CreateBy
    a.CreateByname = apply.CreateByname
    a.CreatePhone = apply.CreatePhone
    return nil
}

func (a *Apply) UpdateStatus() error {
    apply := &models.DeployApp{}
    updateData := map[string]interface{}{
        "status": a.Status,
    }
    if ok := apply.UpdateByFields(updateData, models.QueryParam{
        Where: []models.WhereParam{
            models.WhereParam{
                Field: "id",
                Prepare: a.ID,
            },
        },
    }); !ok {
        return errors.New("update deploy apply status failed")
    }

    return nil
}


func (a *Apply) CheckHaveDeploying() (bool, error) {
    apply := &models.DeployApp{}
    count, ok := apply.Count(models.QueryParam{
        Where: []models.WhereParam{
            models.WhereParam{
                Field: "id",
                Tag: "!=",
                Prepare: a.ID,
            },
            models.WhereParam{
                Field:   "status",
                Prepare: APPLY_STATUS_ING,
            },
            models.WhereParam{
                Field: "created_at",
                Tag: ">=",
                Prepare: int(time.Now().Unix()) - 86400,
            },
        },
    })
    if !ok {
        return false, errors.New("get apply count failed")
    }

    return count == 0, nil
}