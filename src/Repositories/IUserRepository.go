package repositories

import (
	Entity "HotelSystem-LearnGo/Entities"
	"HotelSystem-LearnGo/Models/Requests"
)

type IUserRepository interface {
	Create(user Entity.User) (Entity.User, error)
	Update(user Entity.User) (Entity.User, error)
	Delete(id uint) (string, error)
	FindByEmail(email string) (Entity.User, error)
	FindById(id uint) (Entity.User, error)
	FindByUsername(username string) (Entity.User, error)
	GetAll(req Requests.GetAllRequest) ([]Entity.User, error)
	UnassignRolesFromUser(user Entity.User, roles []Entity.Role) (Entity.User, error)
	AssignRolesToUser(user Entity.User, roles []Entity.Role) (Entity.User, error)
	RemoveAllRolesFromUser(user Entity.User) (Entity.User, error)
	UnassignBuildingsFromUser(user Entity.User, buildings []Entity.Building) (Entity.User, error)
	AssignBuildingsToUser(user Entity.User, buildings []Entity.Building) (Entity.User, error)
	RemoveAllBuildingsFromUser(user Entity.User) (Entity.User, error)
}
