package dto

import "github.com/google/uuid"

type BuildingDto struct {
	Id          uuid.UUID
	Name        string
	Address     string
	City        string
	Province    string
	ZipCode     string
	Description string
}
