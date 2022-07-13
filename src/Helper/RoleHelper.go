package helper

import (
	Entity "HotelSystem-LearnGo/Entities"
	"HotelSystem-LearnGo/Models/Responses"
)

func ToRoleResponse(role Entity.Role) Responses.RoleResponse {

	return Responses.RoleResponse{
		Model:       role.Model,
		Name:        role.Name,
		Description: role.Description,
	}
}

func ToRoleListResponse(roles []Entity.Role) []Responses.RoleResponse {
	var roleResponses []Responses.RoleResponse
	for _, role := range roles {
		roleResponses = append(roleResponses, ToRoleResponse(role))
	}
	return roleResponses
}
