package models

import "testing"

var club Club
var tmpClub Club
var venue Venue

func setupDB() {
	MigrateDB()

	club = Club{
		Name:        "Denny's Club",
		DisplayName: "Denny's Club",
		Status:      false,
	}

	// tmp club with venues
	tmpClub = club

	venue = Venue{
		StreetName: "121, Pearl residency ",
		Locality:   "SVP Nagar",
		State:      "Tamilnadu",
		City:       "Madurai",
		Country:    "India",
		Zipcode:    "625017",
	}

	tmpClub.Venues = []Venue{}
	tmpClub.Venues = append(tmpClub.Venues, venue)
}

func tearDownDB() {
	db.DropTable(&Court{})
	db.DropTable(&Venue{})
	db.DropTable(&Club{})
}

func TestCreateClub(t *testing.T) {
	setupDB()

	err := club.Create()
	if err == nil {
		t.Log("Failed to create club, without Venues , Awesome")
	}

	err = tmpClub.Create()
	if err != nil {
		t.Fail()
		t.Log("Failed to create club ", err)
	}
	t.Log("Test successful, created new club")
	tearDownDB()
}

func TestGetClub(t *testing.T) {

	setupDB()
	//var testClub Club
	tmpClub.Create()

	err := tmpClub.Get(int(tmpClub.ID))
	if err != nil {
		t.Log("Unable to find Club with id ")
		t.Fail()
	}
	tearDownDB()
}

func TestUpdateClub(t *testing.T) {
	setupDB()
	err := tmpClub.Create()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	tmpClub.DisplayName = "cool"
	err = tmpClub.Update(int(tmpClub.ID))
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	tearDownDB()
}

func TestDeleteClub(t *testing.T) {
	setupDB()
	tmpClub.Create()

	err := tmpClub.Delete(int(tmpClub.ID))
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	tearDownDB()
}
