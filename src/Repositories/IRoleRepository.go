package repositories

import (
	Entity "HotelSystem-LearnGo/Entities"
	"HotelSystem-LearnGo/Models/Requests"
)

type IRoleRepository interface {
	Create(role Entity.Role) Entity.Role
	Update(role Entity.Role) Entity.Role
	Delete(id uint) string
	FindByRoleName(roleName string) Entity.Role
	FindById(roleName uint) Entity.Role
	GetAll(req Requests.GetAllRequest) []Entity.Role
}
