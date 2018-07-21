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
		// Communication details
		Phone string `gorm:"not null;index" json:"phone"`
		Email string `gorm:"not null;index" json:"email"`
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
		logrus.Errorf("No venues found in the club %+v\n", club)
		return errors.New("No venues found in the club ")
	}
	err := db.Create(&club).Error
	if err != nil {
		logrus.Errorln("Failed to create new club with values %+v\n", club)
		return err
	}
	return nil
}

// Get a club with id
func (club *Club) Get(id int) error {

	err := db.Model(&club).Where("id=?", id).Preload("Venues").First(&club).Error
	if err != nil {
		logrus.Errorf("Unable to find the club with id  %d ", id)
		return err
	}
	return nil
}

// Update an existing club
func (club *Club) Update(id int) error {

	err := db.Model(&club).Where("id=?", id).Updates(&club).Error

	if err != nil {
		logrus.Errorf("Failed to update clun with id %d and values %+v ", id, club)
		return err
	}

	return nil
}

// Delete an exising club
func (club *Club) Delete(id int) error {
	// transaction
	tx := db.Begin()

	// step1 : query the club and load it
	tx.Model(&club).Where("id=?", id).Preload("Venues").First(&club)
	if tx.Error != nil {
		err := tx.Error
		logrus.Errorf("Failed to find club with id %d for deletion", id)
		tx.Rollback()
		return err
	}
	// step2: update the club
	tx.Model(&club).Where("id=?", id).Update("Status", false)
	if tx.Error != nil {
		err := tx.Error
		logrus.Errorf("Failed to deactivate club with id %d ", id)
		tx.Rollback()
		return err
	}
	// step3: update all venues
	tx.Model(&Venue{}).Where("club_id=?", id).Update("Status", false)
	if tx.Error != nil {
		err := tx.Error
		logrus.Errorf("Failed to deactivate club with id %d ", id)
		tx.Rollback()
		return err
	}
	// step4: iterate to find the venues
	for _, venue := range club.Venues {
		tx.Model(&Court{}).Where("venue_id=?", venue.ID).Update("Status", false)
		if tx.Error != nil {
			err := tx.Error
			logrus.Errorf("Failed to deactivate club with id %d ", id)
			tx.Rollback()
			return err
		}
	}
	// step5 : commit
	tx.Commit()

	return nil
}
