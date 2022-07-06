package Responses

type BuildingResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
	City        string `json:"city"`
	Province    string `json:"province"`
	ZipCode     string `json:"zipcode"`
}
