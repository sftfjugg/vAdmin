// Copyright 2019 syncd Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package common

import (
	"fmt"
	"io/ioutil"
	"vAdmin/module/deploy"
	"vAdmin/tools/command"
	"vAdmin/tools/gofile"
	"vAdmin/tools/gostring"
)

func HookDeploy(applyId int) {
	apply := &deploy.Apply{
		ID: applyId,
	}
	if err := apply.Detail(); err != nil {
		return
	}
	if apply.CommitVersion == "" {
		return
	}

	tmpFile := gostring.JoinStrings("/tmp", "/", gostring.StrRandom(24), ".log")
	if err := gofile.CreateFile(tmpFile, []byte(apply.CommitVersion), 0744); err != nil {
		fmt.Println("apply.CommitVersion tmp file create failed, err[%s], apply_id[%d]", err.Error(), applyId)
		return
	}
	if err := gofile.ReadlineAddHead(tmpFile,"ls"); err != nil {
		fmt.Println("read file change failed, err[%s], apply_id[%d]", err.Error(), applyId)
		return
	}

	var deployCmd string
	data, err := ioutil.ReadFile(tmpFile)
	if err != nil {
		fmt.Println(err)
	} else {
		//fmt.Println(string(data))
		deployCmd = string(data)
	}

	var deployStatus int
	if apply.Status == deploy.APPLY_STATUS_SUCCESS {
		deployStatus = 1
	}

	s := gostring.JoinStrings(
        "#!/bin/bash\n\n",
        "#--------- deploy hook scripts env ---------\n",
        fmt.Sprintf("env_apply_id=%d\n", apply.ID),
		fmt.Sprintf("env_apply_name=%s\n", apply.DemandId),
		fmt.Sprintf("env_deploy_status=%d\n", deployStatus),
		deployCmd,
	)
	hookCommandTaskRun(s, applyId)
}

func hookCommandTaskRun(s string, applyId int) {
	scriptFile := gostring.JoinStrings("/tmp", "/", gostring.StrRandom(24), ".sh")
	if err := gofile.CreateFile(scriptFile, []byte(s), 0744); err != nil {
		fmt.Println("hook script file create failed, err[%s], apply_id[%d]", err.Error(), applyId)
        return
	}
	cmds := []string{
		fmt.Sprintf("/bin/bash -c %s", scriptFile),
		//fmt.Sprintf("rm -f %s", scriptFile),
	}
	task := command.NewTask(cmds, 86400)
	task.Run()
	if err := task.GetError(); err != nil {
		fmt.Println("hook script run failed, err[%s], output[%s], apply_id[%d]", err.Error(), gostring.JsonEncode(task.Result()), applyId)
    }
}