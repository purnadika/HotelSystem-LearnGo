package repositories

import (
	Entity "HotelSystem-LearnGo/Entities"
	"HotelSystem-LearnGo/Models/Requests"
)

type IUserRepository interface {
	Create(user Entity.User) Entity.User
	Update(user Entity.User) Entity.User
	Delete(id uint) string
	FindByEmail(email string) Entity.User
	FindById(id uint) Entity.User
	FindByUsername(username string) Entity.User
	GetAll(req Requests.GetAllRequest) []Entity.User
	UnassignRolesFromUser(user Entity.User, roles []Entity.Role) Entity.User
	AssignRolesToUser(user Entity.User, roles []Entity.Role) Entity.User
	RemoveAllRolesFromUser(user Entity.User) Entity.User
	UnassignBuildingsFromUser(user Entity.User, buildings []Entity.Building) Entity.User
	AssignBuildingsToUser(user Entity.User, buildings []Entity.Building) Entity.User
	RemoveAllBuildingsFromUser(user Entity.User) Entity.User
}
