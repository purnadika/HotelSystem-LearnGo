package repositories

import (
	Entity "HotelSystem-LearnGo/Entities"
	"HotelSystem-LearnGo/Models/Requests"
)

type IBuildingRepository interface {
	Create(building Entity.Building) Entity.Building
	Update(building Entity.Building) Entity.Building
	Delete(id uint) string
	FindByBuildingName(buildingName string) Entity.Building
	FindById(id uint) Entity.Building
	GetAll(req Requests.GetAllRequest) []Entity.Building
}
