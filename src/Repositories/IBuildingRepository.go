package repositories

import (
	Entity "HotelSystem-LearnGo/Entities"
	"HotelSystem-LearnGo/Models/Requests"
)

type IBuildingRepository interface {
	Create(building Entity.Building) (Entity.Building, error)
	Update(building Entity.Building) (Entity.Building, error)
	Delete(id uint) (string, error)
	FindByBuildingName(buildingName string) (Entity.Building, error)
	FindById(id uint) (Entity.Building, error)
	GetAll(req Requests.GetAllRequest) ([]Entity.Building, error)
}
