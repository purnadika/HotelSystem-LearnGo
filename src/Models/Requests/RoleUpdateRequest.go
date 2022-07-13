package Requests

type RoleUpdateRequest struct {
	Id          uint   `validate:"required" json:"id"`
	Name        string `validate:"required,min=1,max=100" json:"name"`
	Description string `validate:"required,min=1,max=255" json:"description"`
}
