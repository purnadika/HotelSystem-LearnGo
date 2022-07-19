package Services

import (
	exceptions "HotelSystem-LearnGo/Exceptions"
	helper "HotelSystem-LearnGo/Helper"
	"HotelSystem-LearnGo/Models/Requests"
	repositories "HotelSystem-LearnGo/Repositories"
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

	building, err = service.BuildingRepository.Create(building)
	helper.PanicIfError(err)
	return building
}

func (service *BuildingService) Update(req Requests.BuildingUpdateRequest) Entity.Building {
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)
	currentBuilding, err := service.BuildingRepository.FindById(req.Id)
	if err != nil {
		exceptions.NewNotFoundError(err.Error())
	}
	building := Entity.Building{
		Model:       currentBuilding.Model,
		Name:        req.Name,
		Description: req.Description,
	}

	building, err = service.BuildingRepository.Update(building)
	helper.PanicIfError(err)
	return building
}

func (service *BuildingService) Delete(id uint) string {
	message, err := service.BuildingRepository.Delete(id)
	helper.PanicIfError(err)
	return message
}

func (service *BuildingService) FindByName(name string) Entity.Building {
	building, err := service.BuildingRepository.FindByBuildingName(name)
	if err != nil {
		exceptions.NewNotFoundError(err.Error())
	}
	return building
}
func (service *BuildingService) FindById(id uint) Entity.Building {
	building, err := service.BuildingRepository.FindById(id)
	if err != nil {
		exceptions.NewNotFoundError(err.Error())
	}
	return building
}

func (service *BuildingService) GetAll(req Requests.GetAllRequest) []Entity.Building {
	reqString := helper.SerializeObject(req)
	log.Println("Request " + string(reqString))
	result, err := service.BuildingRepository.GetAll(req)
	helper.PanicIfError(err)
	return result
}
