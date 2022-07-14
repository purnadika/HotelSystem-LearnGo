package repositories

import (
	"HotelSystem-LearnGo/Database"
	helper "HotelSystem-LearnGo/Helper"
	"HotelSystem-LearnGo/Models/Requests"
	"strconv"

	Entity "HotelSystem-LearnGo/Entities"
)

type UserRepository struct {
}

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}

func (repo UserRepository) Create(user Entity.User) Entity.User {
	Database.Instance.Create(&user)
	return user
}

func (repo UserRepository) Update(user Entity.User) Entity.User {
	result := Database.Instance.Save(&user)
	if result.Error != nil {
		helper.PanicIfError(result.Error)
	}
	return user
}

func (repo UserRepository) Delete(id uint) string {
	var user Entity.User
	result := Database.Instance.Delete(&user, id)
	helper.PanicIfError(result.Error)
	return "User with Id [" + strconv.FormatUint(uint64(id), 5) + "] deleted successfully"
}

func (repo UserRepository) FindByUsername(userName string) Entity.User {
	var user Entity.User
	result := Database.Instance.Where("username = ?", userName).First(&user)
	helper.PanicIfError(result.Error)
	return user
}

func (repo UserRepository) FindById(userId uint) Entity.User {
	var user Entity.User
	result := Database.Instance.First(&user, userId)
	if result.Error != nil {
		helper.PanicIfError(result.Error)
	}
	return user
}

func (repo UserRepository) FindByEmail(email string) Entity.User {
	var user Entity.User
	result := Database.Instance.Where("email = ?", email).First(&user)
	helper.PanicIfError(result.Error)
	return user
}

func (repo UserRepository) GetAll(req Requests.GetAllRequest) []Entity.User {
	var users []Entity.User
	orderType := " asc"
	if req.IsDescending {
		orderType = " desc"
	}
	result := Database.Instance.Limit(req.Take).Offset(req.Skip).Order(req.OrderBy + orderType).Find(&users)
	helper.PanicIfError(result.Error)
	return users
}

func (repo UserRepository) UnassignRolesFromUser(user Entity.User, roles []Entity.Role) Entity.User {
	helper.PanicIfError(
		Database.Instance.Model(&user).Association("Roles").Delete(&roles),
	)
	return user
}
func (repo UserRepository) AssignRolesToUser(user Entity.User, roles []Entity.Role) Entity.User {
	helper.PanicIfError(
		Database.Instance.Model(&user).Association("Roles").Append(&roles),
	)
	return user
}

func (repo UserRepository) RemoveAllRolesFromUser(user Entity.User) Entity.User {
	helper.PanicIfError(
		Database.Instance.Model(&user).Association("Roles").Clear(),
	)
	return user
}

func (repo UserRepository) UnassignBuildingsFromUser(user Entity.User, buildings []Entity.Building) Entity.User {
	helper.PanicIfError(
		Database.Instance.Model(&user).Association("Buildings").Delete(&buildings),
	)
	return user
}
func (repo UserRepository) AssignBuildingsToUser(user Entity.User, buildings []Entity.Building) Entity.User {
	helper.PanicIfError(
		Database.Instance.Model(&user).Association("Buildings").Append(&buildings),
	)
	return user
}

func (repo UserRepository) RemoveAllBuildingsFromUser(user Entity.User) Entity.User {
	helper.PanicIfError(
		Database.Instance.Model(&user).Association("Buildings").Clear(),
	)
	return user
}
