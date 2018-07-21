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

	// bind status - will be replaced by prometheus
	v1.GET("/status", func(c echo.Context) error {
		return nil
	})

	// bind all club apis
	v1.POST("/clubs", createClub)
	v1.GET("/clubs/:id", getClub)
}
