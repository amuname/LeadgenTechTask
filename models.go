package main

import (
	"gorm.io/gorm"
	"time"
)

type Building struct {
	gorm.Model
	Name   string `json:"name"`
	City   string `json:"city"`
	Year   uint   `json:"year"`
	Floors int    `json:"floors"`
}

type CreateBuildingDto struct {
	Name   string `json:"name"`
	City   string `json:"city"`
	Year   uint   `json:"year"`
	Floors int    `json:"floors"`
}

type GetBuildingDto struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Name      string    `json:"name"`
	City      string    `json:"city"`
	Year      uint      `json:"year"`
	Floors    int       `json:"floors"`
}
