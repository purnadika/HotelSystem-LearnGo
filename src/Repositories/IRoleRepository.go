package repositories

import (
	Entity "HotelSystem-LearnGo/Entities"
	"HotelSystem-LearnGo/Models/Requests"
)

type IRoleRepository interface {
	Create(role Entity.Role) (Entity.Role, error)
	Update(role Entity.Role) (Entity.Role, error)
	Delete(id uint) (string, error)
	FindByRoleName(roleName string) (Entity.Role, error)
	FindById(roleName uint) (Entity.Role, error)
	GetAll(req Requests.GetAllRequest) ([]Entity.Role, error)
}
