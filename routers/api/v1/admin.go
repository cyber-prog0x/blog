package v1

import "github.com/gin-gonic/gin"

func AdminPannel(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}
