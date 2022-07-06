package dto

import "github.com/google/uuid"

type UserDto struct {
	Id        uuid.UUID
	Name      string
	Email     string
	Password  string
	Address   string
	Buildings []BuildingDto
	Roles     []RoleDto
}
