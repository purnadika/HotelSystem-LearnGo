package Requests

type UserBuildingsRequest struct {
	UserId      uint   `json:"userid" validate:"required"`
	BuildingIds []uint `json:"buildingids" validate:"required"`
}
