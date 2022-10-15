package router

//InitarticleCheckRoleRouter(r, authMiddleware)
import (
	"github.com/gin-gonic/gin"
	"vAdmin/middleware"
	"vAdmin/pkg/jwtauth"
	"vAdmin/apis/article"
)

//业务 需要认证的路由
func InitarticleCheckRoleRouter(v1 *gin.Engine, authMiddleware *jwtauth.GinJWTMiddleware) {

	v1auth := v1.Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		v1auth.GET("/api/v1/articleList", article.GetArticleList)
		v1auth.GET("/api/v1/article/:articleId", article.GetArticle)
		v1auth.POST("/api/v1/article", article.InsertArticle)
		v1auth.PUT("/api/v1/article", article.UpdateArticle)
		v1auth.DELETE("/api/v1/article/:articleId", article.DeleteArticle)
	}

}