package deploy

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"time"
)

// http升级websocket协议的配置
var WSUpgrader = websocket.Upgrader{
	// 允许所有CORS跨域请求
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 客户端读写消息
type WSMessage struct {
	messageType int
	data []byte
}

// 客户端连接
type WSConnection struct {
	WSSocket *websocket.Conn // 底层websocket
	InChan chan *WSMessage   // 读队列
	OutChan chan *WSMessage  // 写队列

	mutex sync.Mutex	// 避免重复关闭管道
	IsClosed bool
	CloseChan chan byte  // 关闭通知
}

func (wsConn *WSConnection)WSReadLoop() {
	for {
		// 读一个message
		msgType, data, err := wsConn.WSSocket.ReadMessage()
		if err != nil {
			goto error
		}
		req := &WSMessage{
			msgType,
			data,
		}
		// 放入请求队列
		select {
		case wsConn.InChan <- req:
		case <- wsConn.CloseChan:
			goto closed
		}
	}
error:
	wsConn.wsClose()
closed:
}

func (wsConn *WSConnection)WSWriteLoop(ExecStatusChan chan int) {
	var (
		ResultExecChan = make(chan string, 1000)
		ExitExecChan   = make(chan int)
	)
	for {
		select {
		// 取一个应答
		case msg := <- wsConn.OutChan:
			// 写给websocket
			if string(msg.data) == "heartbeat from server" {
				/*
				if err := wsConn.WSSocket.WriteMessage(msg.messageType, msg.data); err != nil {
					goto error
				}*/
				fmt.Println("heartbeat from server")
			} else {
				go Execute(string(msg.data), ResultExecChan, ExitExecChan)

				go func() {
					a := <-ExitExecChan // 没有记录的话会堵塞，会等下次的记录被插入才会放行
					if a == 1 {
						fmt.Println("Execute goroute结束")
					}
					close(ResultExecChan) // 1此循环结束代表1个收集素数的goroute全部完成完毕了，此时才能关闭resultchan，为的是让下面for循环取完resultchan管道中值的时可以正常退出
				}()

				for v := range ResultExecChan {
					//fmt.Println(v)
					msg.data = []byte(v)
					if err := wsConn.WSSocket.WriteMessage(msg.messageType, msg.data); err != nil {
						fmt.Println(err)
						ExecStatusChan <- 3
						goto error
					}
				}
				ExecStatusChan <- 2
			}

		case <- wsConn.CloseChan:
			goto closed
		}
	}
error:
	wsConn.wsClose()
closed:
}

func (wsConn *WSConnection)ProcLoop() {
	// 启动一个gouroutine发送心跳
	go func() {
		for {
			time.Sleep(2 * time.Second)
			if err := wsConn.WSWrite(websocket.TextMessage, []byte("heartbeat from server")); err != nil {
				fmt.Println("heartbeat fail")
				wsConn.wsClose()
				break
			}
		}
	}()

	// 这是一个同步处理模型（只是一个例子），如果希望并行处理可以每个请求一个gorutine，注意控制并发goroutine的数量!!!
	for {
		msg, err := wsConn.wsRead()
		if err != nil {
			fmt.Println("read fail")
			break
		}
		fmt.Println(string(msg.data))
		err = wsConn.WSWrite(msg.messageType, msg.data)
		if err != nil {
			fmt.Println("write fail")
			break
		}
	}
}

func (wsConn *WSConnection)WSWrite(messageType int, data []byte) error {
	select {
	case wsConn.OutChan <- &WSMessage{messageType, data,}:
	case <- wsConn.CloseChan:
		return errors.New("websocket closed")
	}
	return nil
}

func (wsConn *WSConnection)wsRead() (*WSMessage, error) {
	select {
	case msg := <- wsConn.InChan:
		return msg, nil
	case <- wsConn.CloseChan:
	}
	return nil, errors.New("websocket closed")
}

func (wsConn *WSConnection)wsClose() {
	wsConn.WSSocket.Close()

	wsConn.mutex.Lock()
	defer wsConn.mutex.Unlock()
	if !wsConn.IsClosed {
		wsConn.IsClosed = true
		close(wsConn.CloseChan)
	}
}

/*
func wsHandler(resp http.ResponseWriter, req *http.Request) {
	// 应答客户端告知升级连接为websocket
	wsSocket, err := wsUpgrader.Upgrade(resp, req, nil)
	if err != nil {
		return
	}
	wsConn := &wsConnection{
		wsSocket: wsSocket,
		inChan: make(chan *wsMessage, 1000),
		outChan: make(chan *wsMessage, 1000),
		closeChan: make(chan byte),
		isClosed: false,
	}

	// 处理器
	go wsConn.procLoop()
	// 读协程
	go wsConn.wsReadLoop()
	// 写协程
	go wsConn.wsWriteLoop()
}
*/

/*
func main() {
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe("0.0.0.0:7777", nil)
}
*/