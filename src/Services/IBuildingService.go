package Services

import (
	Entity "HotelSystem-LearnGo/Entities"
	"HotelSystem-LearnGo/Models/Requests"
)

type IBuildingService interface {
	Create(req Requests.BuildingCreateRequest) Entity.Building
	Update(req Requests.BuildingUpdateRequest) Entity.Building
	Delete(id uint) string
	FindByName(buildingName string) Entity.Building
	FindById(id uint) Entity.Building
	GetAll(request Requests.GetAllRequest) []Entity.Building
}
