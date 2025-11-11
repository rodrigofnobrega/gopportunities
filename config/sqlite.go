package config

import (
	"os"

	"github.com/rodrigofnobrega/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"
	// Check if the database file exists
	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")
		err = os.MkdirAll("./db", os.ModePerm)

		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)

		if err != nil {
			return nil, err
		}

		file.Close()
	}

	// Create DB and Connect
	db, err := gorm.Open(sqlite.Open("./db/main.db"), &gorm.Config{})

	if err != nil {
		logger.Errf("sqlite opening error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})

	if err != nil {
		logger.Errf("sqlite automigration error: %v", err)
		return nil, err
	}

	return db, nil
}
