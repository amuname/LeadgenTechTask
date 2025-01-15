package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(config *Config) (*gorm.DB, error) {
	dsn := "host=" + config.Host + " user=" + config.User + " password=" + config.Password + " dbname=" + config.DB + " port=" + config.Port + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&Building{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
