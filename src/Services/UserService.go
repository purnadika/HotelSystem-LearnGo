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

	user, err = service.UserRepository.Create(user)
	helper.PanicIfError(err)
	return user
}

func (service *UserService) Update(req Requests.UserUpdateRequest) Entity.User {
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)
	currentUser, err := service.UserRepository.FindById(req.ID)
	helper.PanicIfError(err)
	exceptions.NewNotFoundError("User with Id : [" + helper.UintToString(req.ID) + "] not found")
	user := Entity.User{
		Model:   currentUser.Model,
		Name:    req.Name,
		Email:   req.Email,
		Address: req.Address,
	}

	user, err = service.UserRepository.Update(user)
	helper.PanicIfError(err)
	return user
}

func (service *UserService) Delete(id uint) string {
	message, err := service.UserRepository.Delete(id)
	helper.PanicIfError(err)
	return message
}

func (service *UserService) FindByUsername(username string) Entity.User {
	user, err := service.UserRepository.FindByUsername(username)
	helper.PanicIfError(err)
	return user
}

func (service *UserService) FindByEmail(email string) Entity.User {
	user, err := service.UserRepository.FindByEmail(email)
	helper.PanicIfError(err)
	return user
}

func (service *UserService) FindById(id uint) Entity.User {
	user, err := service.UserRepository.FindById(id)
	helper.PanicIfError(err)
	return user
}

func (service *UserService) GetAll(req Requests.GetAllRequest) []Entity.User {
	reqString := helper.SerializeObject(req)
	log.Println("Request " + string(reqString))
	result, err := service.UserRepository.GetAll(req)
	helper.PanicIfError(err)
	return result
}

func (service *UserService) AssignBuildingsToUser(userId uint, buildingIds []uint) Entity.User {
	user, err := service.UserRepository.FindById(userId)
	if err != nil {
		exceptions.NewNotFoundError(err.Error())
	}
	var buildings []Entity.Building
	for _, buildingId := range buildingIds {
		building := service.BuildingService.FindById(buildingId)
		buildings = append(buildings, building)
	}
	result, err := service.UserRepository.AssignBuildingsToUser(user, buildings)
	return result
}

func (service *UserService) AssignRolesToUser(userId uint, roleIds []uint) Entity.User {
	user, err := service.UserRepository.FindById(userId)
	if err != nil {
		exceptions.NewNotFoundError(err.Error())
	}
	var roles []Entity.Role
	for _, roleId := range roleIds {
		role := service.RoleService.FindById(roleId)
		roles = append(roles, role)
	}
	result, err := service.UserRepository.AssignRolesToUser(user, roles)
	return result
}

func (service *UserService) RemoveRolesFromUser(userId uint, roleIds []uint) Entity.User {
	user, err := service.UserRepository.FindById(userId)
	if err != nil {
		exceptions.NewNotFoundError(err.Error())
	}
	var roles []Entity.Role
	for _, roleId := range roleIds {
		role := service.RoleService.FindById(roleId)
		roles = append(roles, role)
	}
	result, err := service.UserRepository.UnassignRolesFromUser(user, roles)
	return result
}

func (service *UserService) RemoveBuildingsFromUser(userId uint, buildingIds []uint) Entity.User {
	user, err := service.UserRepository.FindById(userId)
	if err != nil {
		exceptions.NewNotFoundError(err.Error())
	}
	var buildings []Entity.Building
	for _, buildingId := range buildingIds {
		building := service.BuildingService.FindById(buildingId)
		buildings = append(buildings, building)
	}
	result, err := service.UserRepository.UnassignBuildingsFromUser(user, buildings)
	return result
}

func (service *UserService) RemoveAllBuildingsFromUser(userId uint) Entity.User {
	user, err := service.UserRepository.FindById(userId)
	if err != nil {
		exceptions.NewNotFoundError(err.Error())
	}
	result, err := service.UserRepository.RemoveAllBuildingsFromUser(user)
	return result
}

func (service *UserService) RemoveAllRolesFromUser(userId uint) Entity.User {
	user, err := service.UserRepository.FindById(userId)
	if err != nil {
		exceptions.NewNotFoundError(err.Error())
	}
	result, err := service.UserRepository.RemoveAllRolesFromUser(user)
	return result
}
