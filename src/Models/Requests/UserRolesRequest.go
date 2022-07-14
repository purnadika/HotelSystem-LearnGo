package Requests

type UserRolesRequest struct {
	UserId  uint   `json:"userid" validate:"required"`
	RoleIds []uint `json:"roleids" validate:"required"`
}
