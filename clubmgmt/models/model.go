package models

import (
	"os"

	"github.com/Sirupsen/logrus"

	// All supported drivers
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// InitDB initiaze the database
func init() {

	// read the database type and URL from the environment variables
	dbtype := os.Getenv("DATABASE_TYPE")
	dburl := os.Getenv("DATABASE_URL")

	//
	sess, err := gorm.Open(dbtype, dburl)
	if err != nil {
		logrus.Errorln("Failed to open databsae connection ", err)
	}
	db = sess

}
