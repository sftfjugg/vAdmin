package deploy

import (
	"fmt"
	"vAdmin/tools/command"
)

const (
	COMMAND_TIMEOUT = 3600
)

type Server struct {
	ID              int
	Addr            string
	User            string
	Port            int
	Cmd          	string
	Key             string
	task            *command.Task
	result          *ServerResult
}

type ServerResult struct {
	ID          int
	TaskResult  []*command.TaskResult
	Status      int
	Error       error
}

func NewServer(srv *Server) {
	srv.result = &ServerResult{
		ID: srv.ID,
		Status: STATUS_INIT,
	}
	srv.task = command.NewTask(
		srv.deployCmd(),
		COMMAND_TIMEOUT,
	)
}

func (srv *Server) Deploy() {
	srv.result.Status = STATUS_ING
	srv.task.Run()
	if err := srv.task.GetError(); err != nil {
		srv.result.Error = err
		srv.result.Status = STATUS_FAILED
	} else {
		srv.result.Status = STATUS_DONE
	}
}

func (srv *Server) Terminate() {
	if srv.result.Status == STATUS_ING {
		srv.task.Terminate()
	}
}

func (srv *Server) Result() *ServerResult {
	srv.result.TaskResult = srv.task.Result()
	return srv.result
}

func (srv *Server) deployCmd() []string {
	var (
		useCustomKey, useSshPort string
	)
	if srv.Key != "" {
		useCustomKey = fmt.Sprintf("-i %s", srv.Key)
	}
	if srv.Port != 0 {
		useSshPort = fmt.Sprintf("-p %d", srv.Port)
	}
	var cmds []string

	if srv.Cmd != "" {
		cmds = append(
			cmds,
			fmt.Sprintf("/usr/bin/env ssh -o StrictHostKeyChecking=no %s %s %s@%s '%s'",
				useCustomKey,
				useSshPort,
				srv.User,
				srv.Addr,
				srv.Cmd,
			),
		)
	}
	return cmds
}