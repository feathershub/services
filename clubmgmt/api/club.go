package api

import (
	"github.com/feathershub/services/clubmgmt/models"
	"github.com/labstack/echo"
)

func init() {

}

func createClub(c echo.Context) error {

	var club models.Club
	//bind the input json to the club
	err := c.Bind(&club)
	if err != nil {
		return c.JSON(400, err)
	}
	// create the club
	err = club.Create()
	if err != nil {
		return c.JSON(400, err)
	}
	return c.JSON(200, &club)
}
