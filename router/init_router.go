package router

import (
	"github.com/gin-gonic/gin"
	"vAdmin/middleware"
	_ "vAdmin/pkg/jwtauth"
	"vAdmin/tools"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	middleware.InitMiddleware(r)
	// the jwt middleware
	authMiddleware, err := middleware.AuthInit()
	tools.HasError(err, "JWT Init Error", 500)

	// 注册系统路由
	InitSysRouter(r, authMiddleware)

	// 注册业务路由
	// TODO: 这里可存放业务路由，里边并无实际路由是有演示代码
	InitAppsRouter(r, authMiddleware)

	//InitarticleCheckRoleRouter(r, authMiddleware)
	//InitreleasecicdCheckRoleRouter(r, authMiddleware)
	return r
}
