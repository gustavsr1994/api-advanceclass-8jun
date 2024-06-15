package main

import (
	"fmt"

	"example/api-advance-class/config"
	"example/api-advance-class/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Welcome to Rest")
	config.ConnectDatabase()
	r := gin.Default()
	r.GET("product", controllers.All)
	r.GET("product/:id", controllers.Index)
	r.POST("product", controllers.Create)
	r.PUT("product", controllers.Update)
	r.DELETE("product/:id", controllers.Delete)
	r.Run()
}
