package dto

import "github.com/google/uuid"

type RoleDto struct {
	Id          uuid.UUID
	Name        string
	Description string
}
