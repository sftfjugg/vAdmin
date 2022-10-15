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

func RegisterClassifyRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	classify := v1.Group("/classify").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		classify.GET("", process.ClassifyList)
		classify.POST("", process.CreateClassify)
		classify.PUT("", process.UpdateClassify)
		classify.DELETE("", process.DeleteClassify)
	}
}
