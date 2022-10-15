package process

/*
  @Author : Rongxin Linghu
*/

import (
	"vAdmin/apis/process"
	"vAdmin/middleware"
	jwt "vAdmin/pkg/jwtauth"

	"github.com/gin-gonic/gin"
)

func RegisterProcessRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	processRouter := v1.Group("/process").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		processRouter.GET("/classify", process.ClassifyProcessList)
		processRouter.GET("", process.ProcessList)
		processRouter.POST("", process.CreateProcess)
		processRouter.PUT("", process.UpdateProcess)
		processRouter.DELETE("", process.DeleteProcess)
		processRouter.GET("/details", process.ProcessDetails)
	}
}
