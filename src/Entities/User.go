package Entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string     `json:"name"`
	Email     string     `gorm:"unique"`
	Username  string     `gorm:"unique"`
	Password  string     `json:"password"`
	Address   string     `json:"address"`
	Buildings []Building `json:"buildings" gorm:"many2many:user_buildings;"`
	Roles     []Role     `json:"roles" gorm:"many2many:user_roles;"`
}
