package router

import (
	"github.com/gin-gonic/gin"
	//"vAdmin/apis/tpl"
	"vAdmin/apis/releasecicd"
	"vAdmin/router/process"
	jwt "vAdmin/pkg/jwtauth"
)



// 路由示例
func InitAppsRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) *gin.Engine {

	// 无需认证的路由
	examplesNoCheckRoleRouter(r)
	// 需要认证的路由
	ReleasecicdCheckRoleRouter(r, authMiddleware)

	DeployAppCheckRoleRouter(r, authMiddleware)

	ProcessCheckRoleRouter(r, authMiddleware)

	return r
}

// 无需认证的路由示例
func examplesNoCheckRoleRouter(r *gin.Engine) {

	v1 := r.Group("/api/v1")
	//v1.GET("/examples/list", examples.apis)
	v1.GET("/releasechart/deploybyday",releasecicd.GetDeployByDay)
	v1.GET("/releasechart/deploybyhour",releasecicd.GetDeployByHour)
}

// 需要认证的路由示例
func ReleasecicdCheckRoleRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	v1:= r.Group("/api/v1")
	registerReleasecicdRouter(v1,authMiddleware)
	registerReleaseridRouter(v1,authMiddleware)
	registerReleasetodoRouter(v1,authMiddleware)
	//registerReleaseChartRouter(v0,authMiddleware)
	registerReleaseStgRouter(v1,authMiddleware)
	//v1auth := v1.Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	//{
	//	v1auth.GET("/examples/list", examples.apis)
	//}
}

func ProcessCheckRoleRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {

	v1 := r.Group("/api/v1")

	// 兼容前后端不分离的情
	// r.GET("/", tpl.Tpl)

	// 流程中心
	process.RegisterClassifyRouter(v1, authMiddleware)
	process.RegisterProcessRouter(v1, authMiddleware)
	process.RegisterTaskRouter(v1, authMiddleware)
	process.RegisterTplRouter(v1, authMiddleware)
	process.RegisterWorkOrderRouter(v1, authMiddleware)

}

// 需要认证的路由示例
func DeployAppCheckRoleRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	v1:= r.Group("/api/v1")
	registerDeployAppRouter(v1,authMiddleware)
}