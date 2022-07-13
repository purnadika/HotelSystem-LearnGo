package repositories

import (
	Entity "HotelSystem-LearnGo/Entities"
	"HotelSystem-LearnGo/Models/Requests"
)

type IBuildingRepository interface {
	Create(building Entity.Building) Entity.Building
	Update(building Entity.Building) Entity.Building
	Delete(id uint) (string, error)
	FindByBuildingName(buildingName string) (Entity.Building, error)
	GetAll(req Requests.GetAllRequest) []Entity.Building
}
