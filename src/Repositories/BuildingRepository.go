package repositories

import (
	"HotelSystem-LearnGo/Database"
	"HotelSystem-LearnGo/Models/Requests"
	"strconv"

	Entity "HotelSystem-LearnGo/Entities"
)

type BuildingRepository struct {
}

func NewBuildingRepository() IBuildingRepository {
	return &BuildingRepository{}
}

func (repo BuildingRepository) Create(building Entity.Building) (Entity.Building, error) {
	result := Database.Instance.Create(&building)
	return building, result.Error
}

func (repo BuildingRepository) Update(building Entity.Building) (Entity.Building, error) {
	result := Database.Instance.Save(&building)
	return building, result.Error
}

func (repo BuildingRepository) Delete(id uint) (string, error) {
	var building Entity.Building
	result := Database.Instance.Delete(&building, id)
	return "Building with Id [" + strconv.FormatUint(uint64(id), 5) + "] deleted successfully", result.Error
}

func (repo BuildingRepository) FindByBuildingName(buildingName string) (Entity.Building, error) {
	var building Entity.Building
	result := Database.Instance.Where("name = ?", buildingName).First(&building)
	return building, result.Error
}

func (repo BuildingRepository) FindById(buildingId uint) (Entity.Building, error) {
	var building Entity.Building
	result := Database.Instance.First(&building, buildingId)
	return building, result.Error
}

func (repo BuildingRepository) GetAll(req Requests.GetAllRequest) ([]Entity.Building, error) {
	var buildings []Entity.Building
	orderType := " asc"
	if req.IsDescending {
		orderType = " desc"
	}
	result := Database.Instance.Limit(req.Take).Offset(req.Skip).Order(req.OrderBy + orderType).Find(&buildings)
	return buildings, result.Error
}
