package Requests

type UserUpdateRequest struct {
	Id          string   `validate:"required" json:"id"`
	Name        string   `validate:"required" json:"name"`
	Email       string   `validate:"required" json:"email"`
	Password    string   `validate:"required" json:"password"`
	RoleIds     []string `validate:"required" json:"roleIds"`
	Address     string   `validate:"required" json:"address"`
	BuildingIDs []string `validate:"required" json:"buildingIds"`
}
