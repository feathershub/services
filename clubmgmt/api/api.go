package api

import (
	"github.com/labstack/echo"
)

func init() {

}

// InitAPI - intializes and binds all API
func InitAPI(router *echo.Echo) {

	// v1
	v1 := router.Group("v1")

	// bind apis
	v1.GET("/status", func(c echo.Context) error {
		return nil
	})
}
