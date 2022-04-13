package models

import (
	"gorm.io/gorm"
)

type WeatherDetail struct {
	gorm.Model  `json:"-"`
	ID          int    `json:"id"`
	WeatherID   int    `json:"weather_id"`
	Main        string `json:"main"`
	Description string `json:"description"`
}
