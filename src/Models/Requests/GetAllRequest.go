package Requests

type GetAllRequest struct {
	Skip         int    `validate:"required" json:"skip"`
	Take         int    `validate:"required" json:"take"`
	OrderBy      string `validate:"required" json:"orderby"`
	IsDescending bool   `validate:"required" json:"isdescending"`
}
