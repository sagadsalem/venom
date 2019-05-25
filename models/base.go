package models

import (
	"github.com/jinzhu/gorm"
	"os"
)

var (
	// db connection
	db *gorm.DB
	// err handling
	err error
)

func init() {
	// gets all the db credentials & create connection
	DBURI := os.Getenv("DBURI")
	db, err = gorm.Open(os.Getenv("DB_DRIVER"), DBURI)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// make migration for all tables
	migrations(db)
}

// migrations functio to auto migrate the database tables
func migrations(db *gorm.DB) {
	db.Debug().AutoMigrate(
		&User{},
	)
}

// GetDB function to create the database connection
func GetDB() *gorm.DB {
	return db
}
