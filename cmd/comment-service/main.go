package main

import (
	"comment-service/config"
	"comment-service/internal/repository"
	"comment-service/internal/rest"
	"comment-service/internal/service"
	"fmt"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"
)

func main() {
	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	configuration := config.NewConfiguration()
	client := repository.NewMongoDBClient(configuration.MongoDBConfig)
	commentRepository := repository.NewMongoDBRepository(client)
	commentService := service.NewCommentService(commentRepository)
	commentController := rest.NewCommentController(commentService)
	commentController.RegisterRoutes(e)
	serverPort := fmt.Sprintf(configuration.AppConfig.ServerConfig.Port)
	e.Logger.Fatal(e.Start(":" + serverPort))
}
