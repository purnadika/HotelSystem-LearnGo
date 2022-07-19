package repositories

import (
	"HotelSystem-LearnGo/Database"
	"HotelSystem-LearnGo/Models/Requests"
	"strconv"

	Entity "HotelSystem-LearnGo/Entities"
)

type UserRepository struct {
}

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}

func (repo UserRepository) Create(user Entity.User) (Entity.User, error) {
	result := Database.Instance.Create(&user)
	return user, result.Error
}

func (repo UserRepository) Update(user Entity.User) (Entity.User, error) {
	result := Database.Instance.Save(&user)
	return user, result.Error
}

func (repo UserRepository) Delete(id uint) (string, error) {
	var user Entity.User
	result := Database.Instance.Delete(&user, id)
	return "User with Id [" + strconv.FormatUint(uint64(id), 5) + "] deleted successfully", result.Error
}

func (repo UserRepository) FindByUsername(userName string) (Entity.User, error) {
	var user Entity.User
	result := Database.Instance.Where("username = ?", userName).First(&user)
	return user, result.Error
}

func (repo UserRepository) FindById(userId uint) (Entity.User, error) {
	var user Entity.User
	result := Database.Instance.First(&user, userId)
	return user, result.Error
}

func (repo UserRepository) FindByEmail(email string) (Entity.User, error) {
	var user Entity.User
	result := Database.Instance.Where("email = ?", email).First(&user)
	return user, result.Error
}

func (repo UserRepository) GetAll(req Requests.GetAllRequest) ([]Entity.User, error) {
	var users []Entity.User
	orderType := " asc"
	if req.IsDescending {
		orderType = " desc"
	}
	result := Database.Instance.Limit(req.Take).Offset(req.Skip).Order(req.OrderBy + orderType).Find(&users)
	return users, result.Error
}

func (repo UserRepository) UnassignRolesFromUser(user Entity.User, roles []Entity.Role) (Entity.User, error) {
	err := Database.Instance.Model(&user).Association("Roles").Delete(&roles)
	return user, err
}
func (repo UserRepository) AssignRolesToUser(user Entity.User, roles []Entity.Role) (Entity.User, error) {
	err := Database.Instance.Model(&user).Association("Roles").Append(&roles)
	return user, err
}

func (repo UserRepository) RemoveAllRolesFromUser(user Entity.User) (Entity.User, error) {
	err := Database.Instance.Model(&user).Association("Roles").Clear()
	return user, err
}

func (repo UserRepository) UnassignBuildingsFromUser(user Entity.User, buildings []Entity.Building) (Entity.User, error) {
	err := Database.Instance.Model(&user).Association("Buildings").Delete(&buildings)
	return user, err
}
func (repo UserRepository) AssignBuildingsToUser(user Entity.User, buildings []Entity.Building) (Entity.User, error) {
	err := Database.Instance.Model(&user).Association("Buildings").Append(&buildings)
	return user, err
}

func (repo UserRepository) RemoveAllBuildingsFromUser(user Entity.User) (Entity.User, error) {
	err := Database.Instance.Model(&user).Association("Buildings").Clear()
	return user, err
}
