package Requests

type BuildingUpdateRequest struct {
	Id          uint   `validate:"required" json:"id"`
	Name        string `validate:"required" json:"name"`
	Address     string `json:"address"`
	Description string `json:"description"`
	City        string `json:"city"`
	Province    string `json:"province"`
	ZipCode     string `json:"zipcode"`
}
