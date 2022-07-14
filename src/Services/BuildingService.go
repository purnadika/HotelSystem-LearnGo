package Services

import (
	helper "HotelSystem-LearnGo/Helper"
	"HotelSystem-LearnGo/Models/Requests"
	repositories "HotelSystem-LearnGo/Repositories"
	"encoding/json"
	"log"

	Entity "HotelSystem-LearnGo/Entities"

	"github.com/go-playground/validator/v10"
)

type BuildingService struct {
	BuildingRepository repositories.IBuildingRepository
	Validate           *validator.Validate
}

func NewBuildingService(repo repositories.IBuildingRepository, validate *validator.Validate) IBuildingService {
	return &BuildingService{
		BuildingRepository: repo,
		Validate:           validate,
	}
}

func (service *BuildingService) Create(req Requests.BuildingCreateRequest) Entity.Building {
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)

	building := Entity.Building{
		Name:        req.Name,
		Description: req.Description,
	}

	building = service.BuildingRepository.Create(building)
	return building
}

func (service *BuildingService) Update(req Requests.BuildingUpdateRequest) Entity.Building {
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)
	currentBuilding := service.BuildingRepository.FindById(req.Id)
	building := Entity.Building{
		Model:       currentBuilding.Model,
		Name:        req.Name,
		Description: req.Description,
	}

	building = service.BuildingRepository.Update(building)
	return building
}

func (service *BuildingService) Delete(id uint) string {
	message := service.BuildingRepository.Delete(id)
	return message
}

func (service *BuildingService) FindByName(name string) Entity.Building {
	building := service.BuildingRepository.FindByBuildingName(name)
	return building
}
func (service *BuildingService) FindById(id uint) Entity.Building {
	building := service.BuildingRepository.FindById(id)
	return building
}

func (service *BuildingService) GetAll(req Requests.GetAllRequest) []Entity.Building {
	reqString, err := json.Marshal(req)
	helper.PanicIfError(err)
	log.Println("Request " + string(reqString))
	return service.BuildingRepository.GetAll(req)
}
