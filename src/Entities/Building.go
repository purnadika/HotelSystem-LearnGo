package Entity

import "gorm.io/gorm"

type Building struct {
	gorm.Model
	Name        string `gorm:"unique"`
	Address     string
	City        string
	Province    string
	ZipCode     string
	Description string
}
