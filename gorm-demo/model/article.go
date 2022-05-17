package model

import (
	"github.com/gookit/color"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Name    string `json:"name"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

func QueryArticleList() (result []Article) {

	db := GetDb()
	db.Table("article").Find(&result)

	return
}
func QueryAddArticle(article Article) (flag bool) {
	db := GetDb()
	color.Red.Println(article)
	db.Table("article").Create(&article)

	return true
}
func QueryUpdateArticle(article Article) (flag bool) {
	db := GetDb()

	color.Red.Println(article)
	db.Table("article").Save(&article)

	return true
}
func DeleteArticle(id int) (flag bool) {
	db := GetDb()
	var article Article
	db.Table("article").Where("id= ?", id).Delete(&article)
	return true
}
func QueryGetArticleById(id int) (result Article) {
	db := GetDb()

	db.Table("article").Where("id=?", id).First(&result)
	color.Blueln(result)
	return
}
