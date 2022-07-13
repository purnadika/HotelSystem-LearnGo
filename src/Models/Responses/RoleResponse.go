package Responses

import (
	"gorm.io/gorm"
)

type RoleResponse struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}
