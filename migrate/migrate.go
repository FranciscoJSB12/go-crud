package main

import (
	"github.com/FranciscoJSB12/go-crud/initializers"
	"github.com/FranciscoJSB12/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
