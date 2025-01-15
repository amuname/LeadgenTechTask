package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server.
func InitApp() {
	config, err := LoadConfig()
	if err != nil {
		log.Fatalf("error load config file: %v", err)
	}
	db, err := InitDB(config)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	r := gin.Default()

	InitRoutes(r, db)

	err = r.Run(":8080")
	if err != nil {
		log.Fatalf("failed to start application: %v", err)
	}
}
