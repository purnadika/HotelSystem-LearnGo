package Services

import (
	helper "HotelSystem-LearnGo/Helper"
	"HotelSystem-LearnGo/Models/Requests"
	"HotelSystem-LearnGo/Models/Responses"
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

func (service *RoleService) Create(req Requests.RoleCreateRequest) Responses.RoleResponse {
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)

	role := Entity.Role{
		Name:        req.Name,
		Description: req.Description,
	}

	role = service.RoleRepository.Create(role)
	return helper.ToRoleResponse(role)
}

func (service *RoleService) Update(req Requests.RoleUpdateRequest) Responses.RoleResponse {
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)
	currentRole := service.RoleRepository.FindById(req.Id)
	role := Entity.Role{
		Model:       currentRole.Model,
		Name:        req.Name,
		Description: req.Description,
	}

	role = service.RoleRepository.Update(role)
	return helper.ToRoleResponse(role)
}

func (service *RoleService) Delete(id uint) string {
	message := service.RoleRepository.Delete(id)
	return message
}

func (service *RoleService) FindByName(name string) Responses.RoleResponse {
	role := service.RoleRepository.FindByRoleName(name)
	return helper.ToRoleResponse(role)
}
func (service *RoleService) FindById(id uint) Responses.RoleResponse {
	role := service.RoleRepository.FindById(id)
	return helper.ToRoleResponse(role)
}

func (service *RoleService) GetAll(req Requests.GetAllRequest) []Responses.RoleResponse {
	reqString, err := json.Marshal(req)
	helper.PanicIfError(err)
	log.Println("Request " + string(reqString))
	return helper.ToRoleListResponse(service.RoleRepository.GetAll(req))
}
