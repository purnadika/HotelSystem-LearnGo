package repositories

import (
	Entity "HotelSystem-LearnGo/Entities"
	"HotelSystem-LearnGo/Models/Requests"
)

type IUserRepository interface {
	Create(user Entity.User) Entity.User
	Update(user Entity.User) Entity.User
	Delete(id uint) (string, error)
	FindByEmail(email string) (Entity.User, error)
	GetAll(req Requests.GetAllRequest) []Entity.User
}
