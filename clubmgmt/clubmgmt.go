package main

import (
	"github.com/feathershub/services/clubmgmt/api"
	"github.com/labstack/echo"
)

func main() {
	router := echo.New()
	api.InitAPI(router)
	router.Start(":8080")
}
