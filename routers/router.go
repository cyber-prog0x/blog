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

	apiIndex := r.Group("/")
	{
		apiIndex.GET("/", v1.BioPage)
		apiIndex.GET("/dash", v1.AdminPannel)

		apiIndex.GET("/index", v1.IndexPage)
	}

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}

	return r
}
