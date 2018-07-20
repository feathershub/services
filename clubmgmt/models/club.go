package models

import (
	"errors"

	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
)

// all models related to club management
type (
	// Club is the model that abstracts
	Club struct {
		gorm.Model
		// Name of the club must be unique and not null
		Name string `gorm:"unique;not null;unique_index;" json:"name"`
		// DisplayName of the club must be unique and not null
		DisplayName string `gorm:"not null" json:"display_name"`
		// Venues
		Venues []Venue `gorm:"foreignKey:ClubID" json:"venues"`
		// Status of the club, either true (active) or false (inactive)
		Status bool `gorm:"default true;index" json:"status"`
	}

	// Venue belongs to the Club and has many Courts
	Venue struct {
		gorm.Model
		// Address of the venue
		StreetName string `gorm:"not null;" json:"street_name"`
		Locality   string `gorm:"not null;" json:"locality"`
		State      string `gorm:";" json:"state"`
		City       string `gorm:"not null;" json:"city"`
		Country    string `gorm:"not null;" json:"country"`
		Zipcode    string `gorm:"not null;index" json:"zipcode"`
		// ClubID
		ClubID uint `gorm:"" json:"club_id"`
		// Courts
		Courts []Court `gorm:"foreignKey:VenueID" json:"courts"`
		// Status of the venue, either true (active) or false (inactive)
		Status bool `gorm:"default true;index" json:"status"`
	}
	// Court belongs to a Venue
	Court struct {
		gorm.Model
		// VenueID
		VenueID uint `gorm:"" json:"venue_id"`
		// Status of the court, either true (active) or false (inactive)
		Status bool `gorm:"default true;index" json:"status"`
	}
)

// all crud operations on club

// Create a new Club
func (club *Club) Create() error {
	if len(club.Venues) == 0 {
		logrus.Errorln("No venues found in the club ", &club)
		return errors.New("No venues found in the club ")
	}
	err := db.Create(&club).Error
	if err != nil {
		logrus.Errorln("Failed to create new club with values ", &club)
		return err
	}
	return nil
}
