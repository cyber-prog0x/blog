package routers

import (
	v1 "blog/routers/api/v1"

	"github.com/gin-gonic/gin"

	"blog/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	// Add static files
	r.Static("/assets", "./statics")
	r.StaticFile("/favicon.ico", "./statics/img/favicon.ico")
	r.LoadHTMLGlob("templates/*")

	index := r.Group("/")
	{
		index.GET("/", v1.BioHandler)
		index.GET("/dash", v1.LoginHandler)
		index.GET("/index", v1.IndexHandler)
	}

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		apiv1.POST("/add/article", v1.AddArticleHandler)
		apiv1.GET("/del/article/:id", v1.DeleteArticleHandler)

	}

	dashBoard := r.Group("/dashboard")
	{
		dashBoard.GET("/index", v1.DashBoardIndexHandler)
		dashBoard.GET("/edit", v1.EditHandler)
	}

	return r
}
