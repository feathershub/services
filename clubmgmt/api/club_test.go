package api

import (
	"testing"

	"github.com/feathershub/services/clubmgmt/models"
	"github.com/h2non/baloo"
)

var client = baloo.New("http://localhost:8080")

var club = &models.Club{
	Name:        "Denny's Club",
	DisplayName: "Denny's Club",
	Status:      false,
}

func TestCreateClubV1(t *testing.T) {
	client.Post("/v1/clubs").JSON(&club)
}
