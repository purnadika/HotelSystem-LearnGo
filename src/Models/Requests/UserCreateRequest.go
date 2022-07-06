package Requests

type UserCreateRequest struct {
	Name        string   `json:"name"`
	Email       string   `json:"description"`
	Password    string   `json:"password"`
	RoleIds     []string `json:"roleIds"`
	Address     string   `json:"address"`
	BuildingIDs []string `json:"buildingIds"`
}
