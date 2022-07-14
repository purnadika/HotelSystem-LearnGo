package Services

import (
	Entity "HotelSystem-LearnGo/Entities"
	"HotelSystem-LearnGo/Models/Requests"
)

type IUserService interface {
	Create(req Requests.UserCreateRequest) Entity.User
	Update(req Requests.UserUpdateRequest) Entity.User
	Delete(id uint) string
	FindByEmail(email string) Entity.User
	FindByUsername(username string) Entity.User
	FindById(id uint) Entity.User
	GetAll(request Requests.GetAllRequest) []Entity.User
	AssignRolesToUser(userId uint, roleIds []uint) Entity.User
	RemoveRolesFromUser(userId uint, roleIds []uint) Entity.User
	AssignBuildingsToUser(userId uint, buildingIds []uint) Entity.User
	RemoveBuildingsFromUser(userId uint, buildingIds []uint) Entity.User
	RemoveAllRolesFromUser(userId uint) Entity.User
	RemoveAllBuildingsFromUser(userId uint) Entity.User
}
