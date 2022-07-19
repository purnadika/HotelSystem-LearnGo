package Services

import (
	exceptions "HotelSystem-LearnGo/Exceptions"
	helper "HotelSystem-LearnGo/Helper"
	"HotelSystem-LearnGo/Models/Requests"
	repositories "HotelSystem-LearnGo/Repositories"
	"encoding/json"
	"log"

	Entity "HotelSystem-LearnGo/Entities"

	"github.com/go-playground/validator/v10"
)

type RoleService struct {
	RoleRepository repositories.IRoleRepository
	Validate       *validator.Validate
}

func NewRoleService(repo repositories.IRoleRepository, validate *validator.Validate) IRoleService {
	return &RoleService{
		RoleRepository: repo,
		Validate:       validate,
	}
}

func (service *RoleService) Create(req Requests.RoleCreateRequest) Entity.Role {
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)

	role := Entity.Role{
		Name:        req.Name,
		Description: req.Description,
	}

	role, err = service.RoleRepository.Create(role)
	helper.PanicIfError(err)
	return role
}

func (service *RoleService) Update(req Requests.RoleUpdateRequest) Entity.Role {
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)
	currentRole, err := service.RoleRepository.FindById(req.Id)
	helper.PanicIfError(err)
	role := Entity.Role{
		Model:       currentRole.Model,
		Name:        req.Name,
		Description: req.Description,
	}

	role, err = service.RoleRepository.Update(role)
	helper.PanicIfError(err)
	return role
}

func (service *RoleService) Delete(id uint) string {
	message, err := service.RoleRepository.Delete(id)
	helper.PanicIfError(err)
	return message
}

func (service *RoleService) FindByName(name string) Entity.Role {
	role, err := service.RoleRepository.FindByRoleName(name)
	if err != nil {
		exceptions.NewNotFoundError(err.Error())
	}
	return role
}
func (service *RoleService) FindById(id uint) Entity.Role {
	role, err := service.RoleRepository.FindById(id)
	if err != nil {
		exceptions.NewNotFoundError(err.Error())
	}
	return role
}

func (service *RoleService) GetAll(req Requests.GetAllRequest) []Entity.Role {
	reqString, err := json.Marshal(req)
	helper.PanicIfError(err)
	log.Println("Request " + string(reqString))
	result, err := service.RoleRepository.GetAll(req)
	helper.PanicIfError(err)
	return result
}
