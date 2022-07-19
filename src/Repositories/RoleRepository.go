package repositories

import (
	"HotelSystem-LearnGo/Database"
	"HotelSystem-LearnGo/Models/Requests"
	"strconv"

	Entity "HotelSystem-LearnGo/Entities"
)

type RoleRepository struct {
}

func NewRoleRepository() IRoleRepository {
	return &RoleRepository{}
}

func (repo RoleRepository) Create(role Entity.Role) (Entity.Role, error) {
	result := Database.Instance.Create(&role)
	return role, result.Error
}

func (repo RoleRepository) Update(role Entity.Role) (Entity.Role, error) {
	result := Database.Instance.Save(&role)
	return role, result.Error
}

func (repo RoleRepository) Delete(id uint) (string, error) {
	var role Entity.Role
	result := Database.Instance.Delete(&role, id)
	return "Role with Id [" + strconv.FormatUint(uint64(id), 5) + "] deleted successfully", result.Error
}

func (repo RoleRepository) FindByRoleName(roleName string) (Entity.Role, error) {
	var role Entity.Role
	result := Database.Instance.Where("name = ?", roleName).First(&role)
	return role, result.Error
}

func (repo RoleRepository) FindById(roleId uint) (Entity.Role, error) {
	var role Entity.Role
	result := Database.Instance.First(&role, roleId)
	return role, result.Error
}

func (repo RoleRepository) GetAll(req Requests.GetAllRequest) ([]Entity.Role, error) {
	var roles []Entity.Role
	orderType := " asc"
	if req.IsDescending {
		orderType = " desc"
	}
	result := Database.Instance.Limit(req.Take).Offset(req.Skip).Order(req.OrderBy + orderType).Find(&roles)
	return roles, result.Error
}
