package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func ConnectToDB() (*gorm.DB, error) {
	dsn := os.Getenv("DB_CONN_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil
}
