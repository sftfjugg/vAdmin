package task

/*
  @Author : Rongxin Linghu
*/

import (
	"context"
	"vAdmin/pkg/task/worker"
)

func Send(classify string, scriptPath string, params string) {
	worker.SendTask(context.Background(), classify, scriptPath, params)
}
