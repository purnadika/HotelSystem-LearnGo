package helper

import (
	dto "HotelSystem-LearnGo/Models/Dto"
	"HotelSystem-LearnGo/Models/Responses"
)

func ToRoleResponse(role dto.RoleDto) Responses.RoleResponse {

	return Responses.RoleResponse{
		Id:          role.Id.String(),
		Name:        role.Name,
		Description: role.Description,
	}
}

func ToRoleListResponse(roles []dto.RoleDto) []Responses.RoleResponse {
	var roleResponses []Responses.RoleResponse
	for _, role := range roles {
		roleResponses = append(roleResponses, ToRoleResponse(role))
	}
	return roleResponses
}
