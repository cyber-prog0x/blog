package v1

import "github.com/gin-gonic/gin"

func IndexPage(c *gin.Context) {
	c.HTML(200, "postlist.html", gin.H{})
}
