package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/go-programming-tour-book/blog-service/internal/routers/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	article := v1.NewArticle()
	tag := v1.NewTag()

	apiv1 := r.Group("/api/v1")
	{
		tg := apiv1.Group("/tags")
		{
			tg.POST("", tag.Create)
			tg.DELETE("/:id", tag.Delete)
			tg.PUT("/:id", tag.Update)
			tg.PATCH("/:id/state", tag.Update)
			tg.GET("", tag.List)
		}

		ag := apiv1.Group("articles")
		{
			ag.POST("", article.Create)
			ag.DELETE("/:id", article.Delete)
			ag.PUT("/:id", article.Update)
			ag.PATCH("/:id/state", article.Update)
			ag.GET("/:id", article.Get)
			ag.GET("", article.List)
		}
	}

	return r
}
