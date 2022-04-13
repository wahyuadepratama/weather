package models

import (
	"gorm.io/gorm"
)

type Weather struct {
	gorm.Model `json:"-"`
	ID         int     `json:"id"`
	Lat        float32 `json:"lat"`
	Lon        float32 `json:"lon"`
	Timezone   string  `json:"timezone"`
	Pressure   int     `json:"pressure"`
	Humidity   int     `json:"humidity"`
	WindSpeed  float32 `json:"wind_speed"`
	CreatedAt  string  `json:"created_at"`
}
