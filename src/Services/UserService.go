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

type UserService struct {
	UserRepository  repositories.IUserRepository
	Validate        *validator.Validate
	RoleService     IRoleService
	BuildingService IBuildingService
}

func NewUserService(repo repositories.IUserRepository, validate *validator.Validate, roleService IRoleService, buildingService IBuildingService) IUserService {
	return &UserService{
		UserRepository:  repo,
		Validate:        validate,
		RoleService:     roleService,
		BuildingService: buildingService,
	}
}

func (service *UserService) Create(req Requests.UserCreateRequest) Entity.User {
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)

	var roles []Entity.Role
	for _, roleId := range req.RoleIds {
		roles = append(roles, service.RoleService.FindById(roleId))
	}

	var buildings []Entity.Building
	for _, buildingId := range req.BuildingIDs {
		buildings = append(buildings, service.BuildingService.FindById(buildingId))
	}

	user := Entity.User{
		Name:      req.Name,
		Email:     req.Email,
		Address:   req.Address,
		Roles:     roles,
		Buildings: buildings,
		Username:  req.Username,
		Password:  req.Password,
	}

	user = service.UserRepository.Create(user)
	return user
}

func (service *UserService) Update(req Requests.UserUpdateRequest) Entity.User {
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)
	currentUser := service.UserRepository.FindById(req.ID)
	user := Entity.User{
		Model:   currentUser.Model,
		Name:    req.Name,
		Email:   req.Email,
		Address: req.Address,
	}

	user = service.UserRepository.Update(user)
	return user
}

func (service *UserService) Delete(id uint) string {
	message := service.UserRepository.Delete(id)
	return message
}

func (service *UserService) FindByUsername(username string) Entity.User {
	user := service.UserRepository.FindByUsername(username)
	return user
}

func (service *UserService) FindByEmail(email string) Entity.User {
	user := service.UserRepository.FindByEmail(email)
	return user
}

func (service *UserService) FindById(id uint) Entity.User {
	user := service.UserRepository.FindById(id)
	return user
}

func (service *UserService) GetAll(req Requests.GetAllRequest) []Entity.User {
	reqString, err := json.Marshal(req)
	helper.PanicIfError(err)
	log.Println("Request " + string(reqString))
	return service.UserRepository.GetAll(req)
}

func (service *UserService) AssignBuildingsToUser(userId uint, buildingIds []uint) Entity.User {
	user := service.UserRepository.FindById(userId)
	var buildings []Entity.Building
	for _, buildingId := range buildingIds {
		building := service.BuildingService.FindById(buildingId)
		buildings = append(buildings, building)
	}
	return service.UserRepository.AssignBuildingsToUser(user, buildings)
}

func (service *UserService) AssignRolesToUser(userId uint, roleIds []uint) Entity.User {
	user := service.UserRepository.FindById(userId)
	var roles []Entity.Role
	for _, roleId := range roleIds {
		role := service.RoleService.FindById(roleId)
		roles = append(roles, role)
	}
	return service.UserRepository.AssignRolesToUser(user, roles)
}

func (service *UserService) RemoveRolesFromUser(userId uint, roleIds []uint) Entity.User {
	user := service.UserRepository.FindById(userId)
	var roles []Entity.Role
	for _, roleId := range roleIds {
		role := service.RoleService.FindById(roleId)
		roles = append(roles, role)
	}
	return service.UserRepository.UnassignRolesFromUser(user, roles)
}

func (service *UserService) RemoveBuildingsFromUser(userId uint, buildingIds []uint) Entity.User {
	user := service.UserRepository.FindById(userId)
	var buildings []Entity.Building
	for _, buildingId := range buildingIds {
		building := service.BuildingService.FindById(buildingId)
		buildings = append(buildings, building)
	}
	return service.UserRepository.UnassignBuildingsFromUser(user, buildings)
}

func (service *UserService) RemoveAllBuildingsFromUser(userId uint) Entity.User {
	user := service.UserRepository.FindById(userId)
	return service.UserRepository.RemoveAllBuildingsFromUser(user)
}

func (service *UserService) RemoveAllRolesFromUser(userId uint) Entity.User {
	user := service.UserRepository.FindById(userId)
	return service.UserRepository.RemoveAllRolesFromUser(user)
}
