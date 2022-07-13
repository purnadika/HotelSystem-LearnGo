package repositories

import (
	"HotelSystem-LearnGo/Database"
	helper "HotelSystem-LearnGo/Helper"
	"HotelSystem-LearnGo/Models/Requests"
	"strconv"

	Entity "HotelSystem-LearnGo/Entities"
)

type RoleRepository struct {
}

func NewRoleRepository() IRoleRepository {
	return &RoleRepository{}
}

func (repo RoleRepository) Create(role Entity.Role) Entity.Role {
	Database.Instance.Create(&role)
	return role
}

func (repo RoleRepository) Update(role Entity.Role) Entity.Role {
	result := Database.Instance.Save(&role)
	if result.Error != nil {
		helper.PanicIfError(result.Error)
	}
	return role
}

func (repo RoleRepository) Delete(id uint) string {
	var role Entity.Role
	result := Database.Instance.Delete(&role, id)
	helper.PanicIfError(result.Error)
	return "Role with Id [" + strconv.FormatUint(uint64(id), 5) + "] deleted successfully"
}

func (repo RoleRepository) FindByRoleName(roleName string) Entity.Role {
	var role Entity.Role
	result := Database.Instance.Where("name = ?", roleName).First(&role)
	helper.PanicIfError(result.Error)
	return role
}

func (repo RoleRepository) FindById(roleId uint) Entity.Role {
	var role Entity.Role
	result := Database.Instance.First(&role, roleId)
	if result.Error != nil {
		helper.PanicIfError(result.Error)
	}
	return role
}

func (repo RoleRepository) GetAll(req Requests.GetAllRequest) []Entity.Role {
	var roles []Entity.Role
	orderType := " asc"
	if req.IsDescending {
		orderType = " desc"
	}
	result := Database.Instance.Limit(req.Take).Offset(req.Skip).Order(req.OrderBy + orderType).Find(&roles)
	helper.PanicIfError(result.Error)
	return roles
}
