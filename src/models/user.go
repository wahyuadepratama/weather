package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"-"`
	ID         string `json:"id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
