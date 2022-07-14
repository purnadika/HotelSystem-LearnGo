package Services

import (
	Entity "HotelSystem-LearnGo/Entities"
	"HotelSystem-LearnGo/Models/Requests"
)

type IRoleService interface {
	Create(req Requests.RoleCreateRequest) Entity.Role
	Update(req Requests.RoleUpdateRequest) Entity.Role
	Delete(id uint) string
	FindByName(roleName string) Entity.Role
	FindById(id uint) Entity.Role
	GetAll(request Requests.GetAllRequest) []Entity.Role
}
