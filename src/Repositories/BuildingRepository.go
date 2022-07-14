package repositories

import (
	"HotelSystem-LearnGo/Database"
	helper "HotelSystem-LearnGo/Helper"
	"HotelSystem-LearnGo/Models/Requests"
	"strconv"

	Entity "HotelSystem-LearnGo/Entities"
)

type BuildingRepository struct {
}

func NewBuildingRepository() IBuildingRepository {
	return &BuildingRepository{}
}

func (repo BuildingRepository) Create(building Entity.Building) Entity.Building {
	Database.Instance.Create(&building)
	return building
}

func (repo BuildingRepository) Update(building Entity.Building) Entity.Building {
	result := Database.Instance.Save(&building)
	if result.Error != nil {
		helper.PanicIfError(result.Error)
	}
	return building
}

func (repo BuildingRepository) Delete(id uint) string {
	var building Entity.Building
	result := Database.Instance.Delete(&building, id)
	helper.PanicIfError(result.Error)
	return "Building with Id [" + strconv.FormatUint(uint64(id), 5) + "] deleted successfully"
}

func (repo BuildingRepository) FindByBuildingName(buildingName string) Entity.Building {
	var building Entity.Building
	result := Database.Instance.Where("name = ?", buildingName).First(&building)
	helper.PanicIfError(result.Error)
	return building
}

func (repo BuildingRepository) FindById(buildingId uint) Entity.Building {
	var building Entity.Building
	result := Database.Instance.First(&building, buildingId)
	if result.Error != nil {
		helper.PanicIfError(result.Error)
	}
	return building
}

func (repo BuildingRepository) GetAll(req Requests.GetAllRequest) []Entity.Building {
	var buildings []Entity.Building
	orderType := " asc"
	if req.IsDescending {
		orderType = " desc"
	}
	result := Database.Instance.Limit(req.Take).Offset(req.Skip).Order(req.OrderBy + orderType).Find(&buildings)
	helper.PanicIfError(result.Error)
	return buildings
}
