package v1

import "github.com/gin-gonic/gin"

func BioPage(c *gin.Context) {
	c.HTML(200, "indexpage.html", gin.H{})
}
