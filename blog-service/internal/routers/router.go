package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-programming-tour-book/blog-service/docs"
	"github.com/go-programming-tour-book/blog-service/internal/middleware"
	v1 "github.com/go-programming-tour-book/blog-service/internal/routers/api/v1"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())
	url := ginSwagger.URL("http://127.0.0.1:8000/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

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
