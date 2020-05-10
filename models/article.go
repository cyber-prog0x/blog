package models

import "fmt"

type Article struct {
	Model

	Tag_id  int    `json:"tag_id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
	State   int    `json:"state"`
}

func AddArticle(tag_id int, title string, desc string, content string, state int) bool {
	db.Create(&Article{
		Tag_id:  tag_id,
		State:   state,
		Title:   title,
		Desc:    desc,
		Content: content,
	})

	return true
}

// 0 published
// 1 draft
func ChangeArticleStatus(postid, status int) bool {
	sql := fmt.Sprintf("update blog_article set state=%d where id=%d", status, postid)
	db.Raw(sql)

	return true

}

func DeleteArticle(postid int) bool {

	var article Article
	article.ID = postid

	db.Delete(&article)

	return true
}
