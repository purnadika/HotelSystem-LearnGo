package Requests

import "gorm.io/gorm"

type UserUpdateRequest struct {
	gorm.Model
	Name        string   `validate:"required" json:"name"`
	Email       string   `validate:"required" json:"email"`
	Password    string   `validate:"required" json:"password"`
	RoleIds     []string `validate:"required" json:"roleIds"`
	Address     string   `validate:"required" json:"address"`
	BuildingIDs []string `validate:"required" json:"buildingIds"`
}
