package main

import (
	"books/controller"
	"books/middlewares"
	"books/repository"
	"books/service"
	"context"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	defer client.Disconnect(context.TODO())

	if err != nil {
		fmt.Println("Something went wrong: " + err.Error())
		return
	}
	
	repo := repository.NewMongoRepository(client)
	userRepo := repository.NewUserRepository(client)
	authenticationService := service.NewAuthenticationService(userRepo)
	authController := controller.NewAuthController(authenticationService)
	service := service.New(repo)
	controller := controller.New(service)

	fmt.Println(userRepo.GetUser("user"))

	router := gin.Default()

	router.Use(cors.New(
		cors.Config{
			AllowOrigins: []string{"http://localhost:8100"},
			AllowMethods: []string{"*"},
			AllowHeaders: []string{"content-type"},
		},
	))

	auth := router.Group("/auth")
	{
		auth.POST("", authController.Authenticate)
	}

	api := router.Group("/api", middlewares.VerifyToken())
	{
		booksApi := api.Group("/books")
		{
			booksApi.GET("", controller.GetBooks)
			booksApi.GET("/:isbn", controller.GetBook)
			booksApi.GET("/notifications", controller.ListenForNotifications)
			booksApi.POST("", controller.SaveBook)
			booksApi.PUT("", controller.UpdateBook)
		}
	}
	router.Run("localhost:5000")
}