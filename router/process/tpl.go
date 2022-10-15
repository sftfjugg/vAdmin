/*
  @Author : Rongxin Linghu
*/

package process

import (
	"vAdmin/apis/process"
	"vAdmin/middleware"

	//"vAdmin/apis/process"
	//"vAdmin/middleware"
	jwt "vAdmin/pkg/jwtauth"

	"github.com/gin-gonic/gin"
)

func RegisterTplRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	tplRouter := v1.Group("/tpl").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		tplRouter.GET("", process.TemplateList)
		tplRouter.POST("", process.CreateTemplate)
		tplRouter.PUT("", process.UpdateTemplate)
		tplRouter.DELETE("", process.DeleteTemplate)
		tplRouter.GET("/details", process.TemplateDetails)
	}
}
