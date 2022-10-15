package ws
import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"vAdmin/pkg/logger"
	"sync"
)

var once sync.Once

type Connection struct {
	WsConn             *websocket.Conn
	InChan             chan []byte
	OutChan            chan []byte
	CloseReadLoopChan  chan struct{}
	CloseWriteLoopChan chan struct{}
}

func InitConnection(wsConn *websocket.Conn) *Connection {
	conn := &Connection{
		WsConn:             wsConn,
		InChan:             make(chan []byte, 1000),
		OutChan:            make(chan []byte, 1000),
		CloseReadLoopChan:  make(chan struct{}, 1),
		CloseWriteLoopChan: make(chan struct{}, 1),
	}

	go conn.readLoop()
	go conn.writeLoop()

	return conn
}

// API
func (c *Connection) ReadMessage() (data []byte, err error) {
	select {
	case data = <-c.InChan:
	case <-c.CloseReadLoopChan:
		c.Close()
		c.StopWriteLoop()
		err = errors.New("connection is closed, stop read msg")
	}
	return
}

func (c *Connection) WriteMessage(data []byte) (err error) {
	select {
	case c.OutChan <- data:
	case <-c.CloseWriteLoopChan:
		c.Close()
		c.StopReadLoop()
		err = errors.New("connection is closed, stop write msg")
	}
	return
}

func (c *Connection) Close() {
	// websocket 的 Close() 方法是线程安全的，是可重入的。是一个特例，websocket其他的方法不是线程安全的。
	once.Do(func() {
		c.WsConn.Close()
		logger.Error("websocket connect closed")
	})
}

func (c *Connection) StopReadLoop() {
	c.CloseReadLoopChan <- struct{}{}
}

func (c *Connection) StopWriteLoop() {
	c.CloseWriteLoopChan <- struct{}{}
}

// 内部实现
func (c *Connection) readLoop() {
	for {
		_, data, err := c.WsConn.ReadMessage()
		if err != nil {
			c.Close()
			logger.Error(fmt.Sprintf("read message from websocket error, err: %s", err))
			c.StopWriteLoop()
			return
		}
		select {
		case c.InChan <- data:
		case <-c.CloseReadLoopChan:
			logger.Info("end websocket read loop")
			return
		}
	}
}

func (c *Connection) writeLoop() {
	for {
		select {
		case data := <-c.OutChan:
			err := c.WsConn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				logger.Error(fmt.Sprintf("write message to websocket error, err: %s", err))
				c.Close()
				c.StopReadLoop()
				return
			}
		case <-c.CloseWriteLoopChan:
			logger.Info("end websocket write loop")
			return
		}
	}
}
