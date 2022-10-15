package router

import (
	"github.com/gin-gonic/gin"
	"vAdmin/apis/releasecicd"
	"vAdmin/middleware"
	jwt "vAdmin/pkg/jwtauth"
)

//业务 需要认证的路由
func registerReleasecicdRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	r := v1.Group("/releasecicd").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("/:id", releasecicd.GetReleasecicd)
		r.POST("", releasecicd.InsertReleasecicd)
		r.PUT("", releasecicd.UpdateReleasecicd)
		r.DELETE("/:id", releasecicd.DeleteReleasecicd)
	}

	l := v1.Group("").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		l.GET("/releasecicdList",releasecicd.GetReleasecicdList)
		l.GET("/pageList",releasecicd.GetPageList)
	}

}

func registerReleaseridRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	l := v1.Group("").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		l.GET("/releaseridList",releasecicd.GetRidList)
	}

}

func registerReleasetodoRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	r := v1.Group("/releasetodo").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.PUT("", releasecicd.UpdateReleasetodo)
	}

	l := v1.Group("").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		l.GET("/releasetodoList",releasecicd.GetReleasetodolist)
	}

}

func registerReleaseStgRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	l := v1.Group("").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		l.GET("/releasestgList",releasecicd.GetReleaseStglist)
	}

}

/*
func registerReleaseChartRouter(v1 *gin.RouterGroup,authMiddleware *jwt.GinJWTMiddleware) {
	c := v1.Group("releasechart").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())

	{
		c.GET("/deploybyday",releasecicd.GetDeployByDay)
		c.GET("/deploybyhour",releasecicd.GetDeployByHour)
	}
}
*/