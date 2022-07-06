package Requests

type RoleCreateRequest struct {
	Id			string 
	Name        string `validate:"required,min=1,max=100" json:"name"`
	Description string `validate:"required,min=1,max=255" json:"description"`
}
