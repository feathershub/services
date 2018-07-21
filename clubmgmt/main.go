package main

import (
	"github.com/feathershub/services/clubmgmt/api"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/swaggo/echo-swagger"

	// for swagger docs
	_ "github.com/swaggo/echo-swagger/example/docs"
)

func main() {
	router := echo.New()

	// initialize logging
	router.Use(middleware.Logger())

	// init swagger
	router.GET("/swaggers/*", echoSwagger.WrapHandler)

	api.InitAPI(router)
	router.Start(":8080")
}
