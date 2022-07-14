package Requests

type UserCreateRequest struct {
	Name        string `validate:"required" json:"name"`
	Email       string `validate:"required" json:"email"`
	Password    string `validate:"required" json:"password"`
	RoleIds     []uint `validate:"required" json:"roleIds"`
	Address     string `validate:"required" json:"address"`
	BuildingIDs []uint `validate:"required" json:"buildingIds"`
	Username    string `validate:"required" json:"username"`
}
