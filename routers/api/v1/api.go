package v1

import (
	"blog/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.HTML(200, "postlist.html", gin.H{})
}

func BioHandler(c *gin.Context) {
	c.HTML(200, "indexpage.html", gin.H{})
}

func LoginHandler(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}

func DashBoardIndexHandler(c *gin.Context) {
	c.HTML(200, "dashboard.html", gin.H{})
}

func EditHandler(c *gin.Context) {
	c.HTML(200, "dashboard_edit_post.html", gin.H{})
}

func AddArticleHandler(c *gin.Context) {

	tag_id := c.PostForm("tag_id")
	title := c.PostForm("title")
	desc := c.PostForm("desc")
	content := c.PostForm("content")

	i_tag_id, err := strconv.Atoi(tag_id)
	if err != nil {
		log.Fatalf("error convert tag_id from string\n")
	}

	ok := models.AddArticle(i_tag_id, title, desc, content, 0)

	if ok {
		log.Printf("add post ok!\n")
		c.JSON(http.StatusOK, gin.H{
			"resp": "succesful",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"resp": "internalservererror",
		})
	}

}

func DeleteArticleHandler(c *gin.Context) {
	id := c.Param("id")

	pid, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalf("can not convert string to int\n")
		c.JSON(http.StatusInternalServerError, gin.H{
			"resp": "fail",
		})
	}

	ok := models.DeleteArticle(pid)
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"resp": "successful",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"resp": "internalservererror",
		})
	}
}
