package ws

import (
	"bufio"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"vAdmin/tools"
	"vAdmin/tools/app"
	"log"
	"net/http"
	"sync"
	"time"
	"vAdmin/module/deploy"
)

// Manager 所有 websocket 信息(并发读写操作)
type Manager struct {
	Group                   map[string]map[string]*Client
	groupCount, clientCount uint
	Lock                    sync.Mutex
	Register, UnRegister    chan *Client
	Message                 chan *MessageData
	GroupMessage            chan *GroupMessageData
	BroadCastMessage        chan *BroadCastMessageData
}

// Client 单个 websocket 信息
type Client struct {
	Id, Group  string
	Context    context.Context
	CancelFunc context.CancelFunc
	Socket     *websocket.Conn
	Message    chan []byte
}

// messageData 单个发送数据信息
type MessageData struct {
	Id, Group string
	Context   context.Context
	Message   []byte
}

// groupMessageData 组广播数据信息
type GroupMessageData struct {
	Group   string
	Message []byte
}

// 广播发送数据信息
type BroadCastMessageData struct {
	Message []byte
}

// 读信息，从 websocket 连接直接读取数据
func (c *Client) Read(cxt context.Context) {
	defer func(cxt context.Context) {
		WebsocketManager.UnRegister <- c
		log.Printf("client [%s] disconnect", c.Id)
		if err := c.Socket.Close(); err != nil {
			log.Printf("client [%s] disconnect err: %s", c.Id, err)
		}
	}(cxt)

	for {
		if cxt.Err() != nil {
			break
		}
		messageType, message, err := c.Socket.ReadMessage()

		if err != nil || messageType == websocket.CloseMessage {
			break
		}
		log.Printf("client [%s] receive message: %s", c.Id, string(message))
		c.Message <- message
	}
}

// 读信息，从 websocket 连接直接读取数据
func (c *Client) Readlog(cxt context.Context, filename string) {
	defer func(cxt context.Context) {
		WebsocketManager.UnRegister <- c
		log.Printf("client [%s] disconnect", c.Id)
		if err := c.Socket.Close(); err != nil {
			log.Printf("client [%s] disconnect err: %s", c.Id, err)
		}
	}(cxt)

	var (
		//ResultExecChan = make(chan string, 1000)
		ExitExecChan   = make(chan int)
	)

	for {
		if cxt.Err() != nil {
			break
		}
		messageType, message, err := c.Socket.ReadMessage()

		if string(message) == "w" {

			go deploy.ExecuteToLog(string(message),filename, ExitExecChan)
			/*
			go deploy.Execute(string(message), ResultExecChan, ExitExecChan)
			go func() {
				a := <-ExitExecChan // 没有记录的话会堵塞，会等下次的记录被插入才会放行
				if a == 1 {
					fmt.Println("Execute goroute结束")
				}
				close(ResultExecChan) // 1此循环结束代表1个收集素数的goroute全部完成完毕了，此时才能关闭resultchan，为的是让下面for循环取完resultchan管道中值的时可以正常退出
			}()
			for v := range ResultExecChan {
				//fmt.Println(v)
				data := []byte(v)
				if err := c.Socket.WriteMessage(websocket.CloseMessage, data); err != nil {
					log.Printf("client [%s] writemessage err: %s", c.Id, err)
					return
				}
				log.Printf("client [%s] write message: %s", c.Id, string(data))
			}*/
		}

		if err != nil || messageType == websocket.CloseMessage {
			break
		}
		log.Printf("client [%s] receive message: %s", c.Id, string(message))
		c.Message <- message
	}
}

// 写信息，从 channel 变量 Send 中读取数据写入 websocket 连接
func (c *Client) WriteTaskID(cxt context.Context, taskid string) {
	defer func(cxt context.Context) {
		log.Printf("client [%s] disconnect", c.Id)
		if err := c.Socket.Close(); err != nil {
			log.Printf("client [%s] disconnect err: %s", c.Id, err)
		}
	}(cxt)
	sysType := runtime.GOOS
	for {
		if cxt.Err() != nil {
			break
		}
		select {
		case message, ok := <-c.Message:
			if !ok {
				_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if string(message)[0:4] == "ping" {
				//cmd := exec.Command("/bin/bash", "-c", string(message))
				//cmd := exec.Command("cmd", "/C", "dir D: /s/b")
				switch sysType {
				case "linux":
					cmd := exec.Command("/bin/bash", "-c", string(message))
					fmt.Println(cmd)
					//fmt.Println(cmd.Process.Pid) // 打印进程pid
				case "windows":
					cmd := exec.Command("cmd", "/C", string(message))
					fmt.Println(cmd)
					//cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

					stdout, err := cmd.StdoutPipe()
					if err != nil {
						fmt.Println(err)
					}
					//标准错误也输出到标准输出里
					cmd.Stderr = cmd.Stdout

					if err := cmd.Start(); err != nil {
						log.Fatal(err)
					}
					fmt.Println("任务进程Pid--->", cmd.Process.Pid)
					filename := "temp/logs/job/" + taskid + ".pid"
					//ioutil.WriteFile如果文件存在会清空文件然后写入
					/*
					if err2 := ioutil.WriteFile(filename, []byte(string(cmd.Process.Pid)), 0740); err2 != nil {
						fmt.Printf("ioutil.WriteFile failure, err=[%v]\n", err2)
					}*/
					f, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0740)
					f.Write([]byte(string(cmd.Process.Pid)))
					f.Close()
					in := bufio.NewScanner(stdout)
					/*带缓冲的 IO 标准库 一直是 Go 中优化读写操作的利器。对于写操作来说，在被发送到 socket 或硬盘之前，
					IO 缓冲区 提供了一个临时存储区来存放数据，缓冲区存储的数据达到一定容量后才会被"释放"出来进行下一步存储，
					这种方式大大减少了写操作或是最终的系统调用被触发的次数，
					这无疑会在频繁使用系统资源的时候节省下巨大的系统开销。而对于读操作来说，
					缓冲 IO 意味着每次操作能够读取更多的数据，既减少了系统调用的次数，又通过以块为单位读取硬盘数据来更高效地使用底层硬件。
					bufio 包中的 Scanner 扫描器模块，它的主要作用是把数据流分割成一个个标记并除去它们之间的空格。*/
					for in.Scan() {
						//fmt.Println(in.Text())
						switch sysType {
						case "linux":
							cmdRe := ConvertByte2String(in.Bytes(), "UTF8")
							fmt.Println(cmdRe)
							// 信息推送 WriteMessage
							message = []byte(cmdRe)
							log.Printf("client [%s] write message: %s", c.Id, string(message))
							err := c.Socket.WriteMessage(websocket.TextMessage, message)
							if err != nil {
								log.Printf("client [%s] writemessage err: %s", c.Id, err)
							}
						case "windows":
							cmdRe := ConvertByte2String(in.Bytes(), "GB18030")
							fmt.Println(cmdRe)

							// 信息推送 WriteMessage
							message = []byte(cmdRe)
							log.Printf("client [%s] write message: %s", c.Id, string(message))
							err := c.Socket.WriteMessage(websocket.TextMessage, message)
							if err != nil {
								log.Printf("client [%s] writemessage err: %s", c.Id, err)
							}
						}
					}

					fmt.Println("任务进程结束中...", cmd.Process.Pid)
					if err := cmd.Wait(); err != nil {
						fmt.Println(err)
					}
				}
			}else if string(message)[0:8] == "KillTask" {
				pid, err := ioutil.ReadFile("temp/logs/job/" +  string(message)[8:] + ".pid")
				if err != nil{
					fmt.Println("Read file err =", err)
				}
				fmt.Printf("强制杀进程：%d", pid)
				switch sysType {
				case "linux":
					CMD := "kill" + " " + "-9" + " " + string(pid)
					cmd := exec.Command("/bin/bash", "-c", CMD)
					fmt.Println(cmd)
					//fmt.Println(cmd.Process.Pid) // 打印进程pid
				case "windows":
					CMD := "taskkill" + " " + "/pid" + " " + string(pid)
					cmd := exec.Command("cmd", "/C", CMD)
					fmt.Println(cmd)
					//cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
				}
			}else {
				switch sysType {
				case "linux":
					cmdRe := ConvertByte2String(message, "UTF8")
					fmt.Println(cmdRe)
					message = []byte(cmdRe)
					log.Printf("client [%s] write message: %s", c.Id, string(message))
					err := c.Socket.WriteMessage(websocket.TextMessage, message)
					if err != nil {
						log.Printf("client [%s] writemessage err: %s", c.Id, err)
					}
				case "windows":
					cmdRe := ConvertByte2String(message, "UTF8")
					fmt.Println(cmdRe)
					message = []byte(cmdRe)
					log.Printf("client [%s] write message: %s", c.Id, string(message))
					err := c.Socket.WriteMessage(websocket.TextMessage, message)
					if err != nil {
						log.Printf("client [%s] writemessage err: %s", c.Id, err)
					}
				}
			}
		case _ = <-c.Context.Done():
			break
		}
	}
}

// 写信息，从 channel 变量 Send 中读取数据写入 websocket 连接
func (c *Client) Write(cxt context.Context) {
	defer func(cxt context.Context) {
		log.Printf("client [%s] disconnect", c.Id)
		if err := c.Socket.Close(); err != nil {
			log.Printf("client [%s] disconnect err: %s", c.Id, err)
		}
	}(cxt)
	sysType := runtime.GOOS
	for {
		if cxt.Err() != nil {
			break
		}
		select {
		case message, ok := <-c.Message:
			if !ok {
				_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if string(message)[0:4] == "ping" {
				//cmd := exec.Command("/bin/bash", "-c", string(message))
				//cmd := exec.Command("cmd", "/C", "dir D: /s/b")
				switch sysType {
				case "linux":
					cmd := exec.Command("/bin/bash", "-c", string(message))
					fmt.Println(cmd)
					//fmt.Println(cmd.Process.Pid) // 打印进程pid
				case "windows":
					cmd := exec.Command("cmd", "/C", string(message))
					fmt.Println(cmd)
					//cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

				stdout, err := cmd.StdoutPipe()
				if err != nil {
					fmt.Println(err)
				}
				//标准错误也输出到标准输出里
				cmd.Stderr = cmd.Stdout

				if err := cmd.Start(); err != nil {
					log.Fatal(err)
				}
				fmt.Println("任务进程Pid--->", cmd.Process.Pid)
				in := bufio.NewScanner(stdout)
					/*带缓冲的 IO 标准库 一直是 Go 中优化读写操作的利器。对于写操作来说，在被发送到 socket 或硬盘之前，
					IO 缓冲区 提供了一个临时存储区来存放数据，缓冲区存储的数据达到一定容量后才会被"释放"出来进行下一步存储，
					这种方式大大减少了写操作或是最终的系统调用被触发的次数，
					这无疑会在频繁使用系统资源的时候节省下巨大的系统开销。而对于读操作来说，
					缓冲 IO 意味着每次操作能够读取更多的数据，既减少了系统调用的次数，又通过以块为单位读取硬盘数据来更高效地使用底层硬件。
					bufio 包中的 Scanner 扫描器模块，它的主要作用是把数据流分割成一个个标记并除去它们之间的空格。*/
				for in.Scan() {
					//fmt.Println(in.Text())
					switch sysType {
					case "linux":
						cmdRe := ConvertByte2String(in.Bytes(), "UTF8")
						fmt.Println(cmdRe)
						// 信息推送 WriteMessage
						message = []byte(cmdRe)
						log.Printf("client [%s] write message: %s", c.Id, string(message))
						err := c.Socket.WriteMessage(websocket.TextMessage, message)
						if err != nil {
							log.Printf("client [%s] writemessage err: %s", c.Id, err)
						}
					case "windows":
						cmdRe := ConvertByte2String(in.Bytes(), "GB18030")
						fmt.Println(cmdRe)

						// 信息推送 WriteMessage
						message = []byte(cmdRe)
						log.Printf("client [%s] write message: %s", c.Id, string(message))
						err := c.Socket.WriteMessage(websocket.TextMessage, message)
						if err != nil {
							log.Printf("client [%s] writemessage err: %s", c.Id, err)
						}
					}
				}

				fmt.Println("任务进程结束中...", cmd.Process.Pid)
				if err := cmd.Wait(); err != nil {
					fmt.Println(err)
				}
				}
			}else {
				switch sysType {
				case "linux":
					cmdRe := ConvertByte2String(message, "UTF8")
					fmt.Println(cmdRe)
					message = []byte(cmdRe)
					log.Printf("client [%s] write message: %s", c.Id, string(message))
					err := c.Socket.WriteMessage(websocket.TextMessage, message)
					if err != nil {
						log.Printf("client [%s] writemessage err: %s", c.Id, err)
					}
				case "windows":
					cmdRe := ConvertByte2String(message, "UTF8")
					fmt.Println(cmdRe)
					message = []byte(cmdRe)
					log.Printf("client [%s] write message: %s", c.Id, string(message))
					err := c.Socket.WriteMessage(websocket.TextMessage, message)
					if err != nil {
						log.Printf("client [%s] writemessage err: %s", c.Id, err)
					}
				}
			}
		case _ = <-c.Context.Done():
			break
		}
	}
}

// 启动 websocket 管理器
func (manager *Manager) Start() {
	log.Printf("websocket manage start")
	for {
		select {
		// 注册
		case client := <-manager.Register:
			log.Printf("client [%s] connect", client.Id)
			log.Printf("register client [%s] to group [%s]", client.Id, client.Group)

			manager.Lock.Lock()
			if manager.Group[client.Group] == nil {
				manager.Group[client.Group] = make(map[string]*Client)
				manager.groupCount += 1
			}
			manager.Group[client.Group][client.Id] = client
			manager.clientCount += 1
			manager.Lock.Unlock()

		// 注销
		case client := <-manager.UnRegister:
			log.Printf("unregister client [%s] from group [%s]", client.Id, client.Group)
			manager.Lock.Lock()
			if mGroup, ok := manager.Group[client.Group]; ok {
				if mClient, ok := mGroup[client.Id]; ok {
					close(mClient.Message)
					delete(mGroup, client.Id)
					manager.clientCount -= 1
					if len(mGroup) == 0 {
						//log.Printf("delete empty group [%s]", client.Group)
						delete(manager.Group, client.Group)
						manager.groupCount -= 1
					}
					mClient.CancelFunc()
				}
			}
			manager.Lock.Unlock()

			// 发送广播数据到某个组的 channel 变量 Send 中
			//case data := <-manager.boardCast:
			//	if groupMap, ok := manager.wsGroup[data.GroupId]; ok {
			//		for _, conn := range groupMap {
			//			conn.Send <- data.Data
			//		}
			//	}
		}
	}
}

// 处理单个 client 发送数据
func (manager *Manager) SendService() {
	for {
		select {
		case data := <-manager.Message:
			if groupMap, ok := manager.Group[data.Group]; ok {
				if conn, ok := groupMap[data.Id]; ok {
					conn.Message <- data.Message
				}
			}
		}
	}
}

// 处理 group 广播数据
func (manager *Manager) SendGroupService() {
	for {
		select {
		// 发送广播数据到某个组的 channel 变量 Send 中
		case data := <-manager.GroupMessage:
			if groupMap, ok := manager.Group[data.Group]; ok {
				for _, conn := range groupMap {
					conn.Message <- data.Message
				}
			}
		}
	}
}

// 处理广播数据
func (manager *Manager) SendAllService() {
	for {
		select {
		case data := <-manager.BroadCastMessage:
			for _, v := range manager.Group {
				for _, conn := range v {
					conn.Message <- data.Message
				}
			}
		}
	}
}
type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)
//对字符进行转码
func ConvertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}

// 向指定的 client 发送数据
func (manager *Manager) Send(cxt context.Context, id string, group string, message []byte) {
	data := &MessageData{
		Id:      id,
		Context: cxt,
		Group:   group,
		Message: message,
	}
	manager.Message <- data
}

// 向指定的 Group 广播
func (manager *Manager) SendGroup(group string, message []byte) {
	data := &GroupMessageData{
		Group:   group,
		Message: message,
	}
	manager.GroupMessage <- data
}

// 广播
func (manager *Manager) SendAll(message []byte) {
	data := &BroadCastMessageData{
		Message: message,
	}
	manager.BroadCastMessage <- data
}

// 注册
func (manager *Manager) RegisterClient(client *Client) {
	manager.Register <- client
}

// 注销
func (manager *Manager) UnRegisterClient(client *Client) {
	manager.UnRegister <- client
}

// 当前组个数
func (manager *Manager) LenGroup() uint {
	return manager.groupCount
}

// 当前连接个数
func (manager *Manager) LenClient() uint {
	return manager.clientCount
}

// 获取 wsManager 管理器信息
func (manager *Manager) Info() map[string]interface{} {
	managerInfo := make(map[string]interface{})
	managerInfo["groupLen"] = manager.LenGroup()
	managerInfo["clientLen"] = manager.LenClient()
	managerInfo["chanRegisterLen"] = len(manager.Register)
	managerInfo["chanUnregisterLen"] = len(manager.UnRegister)
	managerInfo["chanMessageLen"] = len(manager.Message)
	managerInfo["chanGroupMessageLen"] = len(manager.GroupMessage)
	managerInfo["chanBroadCastMessageLen"] = len(manager.BroadCastMessage)
	return managerInfo
}

// 初始化 wsManager 管理器
var WebsocketManager = Manager{
	Group:            make(map[string]map[string]*Client),
	Register:         make(chan *Client, 128),
	UnRegister:       make(chan *Client, 128),
	GroupMessage:     make(chan *GroupMessageData, 128),
	Message:          make(chan *MessageData, 128),
	BroadCastMessage: make(chan *BroadCastMessageData, 128),
	groupCount:       0,
	clientCount:      0,
}

// gin 处理 websocket handler
func (manager *Manager) WsDeployClient(c *gin.Context) {

	ctx, cancel := context.WithCancel(context.Background())

	upGrader := websocket.Upgrader{
		// cross origin domain
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		// 处理 Sec-WebSocket-Protocol Header
		Subprotocols: []string{c.GetHeader("Sec-WebSocket-Protocol")},
	}

	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("websocket connect error: %s", c.Param("channel"))
		return
	}

	fmt.Println("token: ", c.Query("token"))

	client := &Client{
		Id:         c.Param("id"),
		Group:      c.Param("channel"),
		Context:    ctx,
		CancelFunc: cancel,
		Socket:     conn,
		Message:    make(chan []byte, 1024),
	}

	manager.RegisterClient(client)
	go client.Read(ctx)
	go client.Write(ctx)
	time.Sleep(time.Second * 15)

	fmt.Println("demand_id:",c.Param("demandid"))
	//filename :=  fmt.Sprintf("temp/logs/job/%s.log", c.Param("demandid"))

	tools.FileMonitoringById(ctx, "temp/logs/job/xxxx.log", c.Param("id"), c.Param("channel"), SendOne)
}

// gin 处理 websocket handler
func (manager *Manager) WsClient(c *gin.Context) {

	ctx, cancel := context.WithCancel(context.Background())

	upGrader := websocket.Upgrader{
		// cross origin domain
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		// 处理 Sec-WebSocket-Protocol Header
		Subprotocols: []string{c.GetHeader("Sec-WebSocket-Protocol")},
	}

	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("websocket connect error: %s", c.Param("channel"))
		return
	}

	fmt.Println("token: ", c.Query("token"))

	client := &Client{
		Id:         c.Param("id"),
		Group:      c.Param("channel"),
		Context:    ctx,
		CancelFunc: cancel,
		Socket:     conn,
		Message:    make(chan []byte, 1024),
	}

	filename :=  fmt.Sprintf("temp/logs/job/%s.log", c.Param("id"))
	file,er:=os.Open(filename)
	if er!=nil && os.IsNotExist(er){
		file,_ = os.Create(filename)
		file.Close()
	}
	manager.RegisterClient(client)

	go client.Read(ctx)
	//go client.Readlog(ctx,filename)
	go client.WriteTaskID(ctx,string(c.Param("id")))
	time.Sleep(time.Second * 5)
	tools.FileMonitoringById(ctx, filename, c.Param("id"), c.Param("channel"), SendOne)

}

func (manager *Manager) UnWsClient(c *gin.Context) {
	id := c.Param("id")
	group := c.Param("channel")
	WsLogout(id, group)
	app.OK(c, "ws close success", "success")
}

func SendGroup(msg []byte) {
	//WebsocketManager.SendGroup("leffss", []byte("{\"code\":200,\"data\":"+string(msg)+"}"))
	WebsocketManager.SendGroup("log", []byte("{\"code\":200,\"data\":"+string(msg)+"}"))
	fmt.Println(WebsocketManager.Info())
}

func SendAll(msg []byte) {
	WebsocketManager.SendAll([]byte("{\"code\":200,\"data\":" + string(msg) + "}"))
	fmt.Println(WebsocketManager.Info())
}

func SendOne(ctx context.Context, id string, group string, msg []byte) {
	WebsocketManager.Send(ctx, id, group, []byte("{\"code\":200,\"data\":"+string(msg)+"}"))
	fmt.Println(WebsocketManager.Info())
}

func WsLogout(id string, group string) {
	WebsocketManager.UnRegisterClient(&Client{Id: id, Group: group})
	fmt.Println(WebsocketManager.Info())
}
