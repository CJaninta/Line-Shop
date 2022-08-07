package main

import (
	_ "assignment/config"
	"assignment/controllers"
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ctx         context.Context
	mongoclient *mongo.Client
)

func main() {
	fmt.Println("Starting app...")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())

	route := router.Group("/api")
	{
		route.POST("/books", controllers.CreateBook)
		route.GET("/books", controllers.GetBook)

		route.GET("/search", controllers.SearchBook)

		route.POST("/promotions", controllers.CreatePromotion)
		route.GET("/promotions", controllers.GetPromotion)

		route.POST("/orders", controllers.CreateOrder)
		route.POST("/orders/checkout", controllers.CheckoutOrder)
		route.GET("/orders/summary", controllers.Summary)
	}

	if err := router.Run(":" + port); err != nil {
		fmt.Print("Server error: ", err)
	}

	defer mongoclient.Disconnect(ctx)

}
