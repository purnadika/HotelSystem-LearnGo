package Services

import (
	"HotelSystem-LearnGo/Models/Requests"
	"HotelSystem-LearnGo/Models/Responses"
)

type IRoleService interface {
	Create(req Requests.RoleCreateRequest) Responses.RoleResponse
	Update(req Requests.RoleUpdateRequest) Responses.RoleResponse
	Delete(id uint) string
	FindByName(roleName string) Responses.RoleResponse
	FindById(id uint) Responses.RoleResponse
	GetAll(request Requests.GetAllRequest) []Responses.RoleResponse
}
