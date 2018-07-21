package api

import (
	"testing"

	"github.com/feathershub/services/clubmgmt/models"
	"github.com/h2non/baloo"
)

var client = baloo.New("http://localhost:8080")

var club models.Club
var tmpClub models.Club
var venue models.Venue

func setup() {

	club = models.Club{
		Name:        "Denny's Club",
		DisplayName: "Denny's Club",
		Status:      false,
	}

	// tmp club with venues
	tmpClub = club

	venue = models.Venue{
		StreetName: "121, Pearl residency ",
		Locality:   "SVP Nagar",
		State:      "Tamilnadu",
		City:       "Madurai",
		Country:    "India",
		Zipcode:    "625017",
	}

	tmpClub.Venues = []models.Venue{}
	tmpClub.Venues = append(tmpClub.Venues, venue)
}

func TestCreateClubV1(t *testing.T) {
	err := client.Post("/v1/clubs").Expect(t).Status(400).Type("json").JSON(club).Done()

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	err = client.Post("/v1/clubs").Expect(t).Status(200).Type("json").JSON(tmpClub).Done()

	if err != nil {
		t.Log(err)
		t.Fail()
	}
}
