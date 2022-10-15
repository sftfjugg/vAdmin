package deploy

import (
	"github.com/gorilla/websocket"
	//"github.com/gin-gonic/gin"
	"net/http"
)

var (
	upGrader websocket.Upgrader
	wsUpdateInterval int64
)

func InitWs() {

	/*
	router := gin.Default()
	router.GET("/api/v1/deployws", DeployStart)
	router.Run(":3000")
	*/
	upGrader = websocket.Upgrader{
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	wsUpdateInterval = 10
}
