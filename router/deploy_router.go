package router

import (
	"github.com/gin-gonic/gin"
	"vAdmin/apis/deploy"
	"vAdmin/middleware"
	jwt "vAdmin/pkg/jwtauth"
)

//业务 需要认证的路由

func registerDeployAppRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	r := v1.Group("/deploy").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("/:id", deploy.GetDeployApp)
		r.POST("", deploy.InsertDeployApp)
		r.PUT("", deploy.UpdateDeployApp)
	}

	l := v1.Group("").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		l.GET("/applyList", deploy.GetDeployAppList)
		l.GET("/rollList", deploy.GetRollbackAppList)
		l.GET("/deploystatus", deploy.DeployStatus)
		l.GET("/deploystart", deploy.DeployStart)
		//l.GET("/killtask", deploy.KillTask)
	}

}