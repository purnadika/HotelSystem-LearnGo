package Services

import (
	"HotelSystem-LearnGo/Models/Requests"
	"HotelSystem-LearnGo/Models/Responses"
)

type IBuildingService interface {
	Create(req Requests.BuildingCreateRequest) Responses.BuildingResponse
	Update(req Requests.BuildingUpdateRequest) Responses.BuildingResponse
	Delete(uuidString string) string
	FindByBuildingName(buildingName string) Responses.BuildingResponse
	GetAll(request Requests.GetAllRequest) []Responses.BuildingResponse
}
