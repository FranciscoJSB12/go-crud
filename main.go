package main

import (
	"github.com/FranciscoJSB12/go-crud/controllers"
	"github.com/FranciscoJSB12/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()

	router.GET("/posts", controllers.GetPosts)
	router.GET("/posts/:id", controllers.GetOnePost)
	router.POST("/posts", controllers.CreatePost)
	router.PUT("/posts/:id", controllers.UpdatePost)
	router.DELETE("/posts/:id", controllers.DeletePost)

	router.Run()
}
