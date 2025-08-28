package main

import (
	"os"

	"github.com/kun1ts4/checklist/api/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/kun1ts4/checklist/api/docs"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
}

// @title           Checklist API
// @version         0.1.1
// @host            localhost:8080
// @BasePath        /
func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.Secure())
	e.Use(middleware.CORS())
	
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/create", handler.CreateTask)
	e.GET("/list", handler.GetTasks)
	e.DELETE("/delete", handler.DeleteTask)
	e.PUT("/done", handler.CompleteTask)

	e.Logger.Fatal(e.Start(":8080"))
}
