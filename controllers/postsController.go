package controllers

import (
	"github.com/FranciscoJSB12/go-crud/initializers"
	"github.com/FranciscoJSB12/go-crud/models"
	"github.com/gin-gonic/gin"
)

func GetPosts(ctx *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	ctx.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetOnePost(ctx *gin.Context) {
	id := ctx.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	ctx.JSON(200, gin.H{
		"posts": post,
	})
}

func CreatePost(ctx *gin.Context) {

	var requestBody struct {
		Title string
		Body  string
	}

	ctx.Bind(&requestBody)

	post := models.Post{Title: requestBody.Title, Body: requestBody.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		ctx.Status(400)
		return
	}

	ctx.JSON(200, gin.H{
		"post": post,
	})
}

func UpdatePost(ctx *gin.Context) {

	id := ctx.Param("id")

	var requestBody struct {
		Title string
		Body  string
	}

	ctx.Bind(&requestBody)

	var post models.Post

	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{
		Title: requestBody.Title,
		Body:  requestBody.Body,
	})

	ctx.JSON(200, gin.H{
		"post": post,
	})
}

func DeletePost(ctx *gin.Context) {
	id := ctx.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	ctx.Status(204)
}
