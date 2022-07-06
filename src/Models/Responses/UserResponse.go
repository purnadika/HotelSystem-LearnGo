package Responses

type UserResponse struct {
	Id        string      `json:"id"`
	Name      string      `json:"name"`
	Email     string      `json:"description"`
	Password  string      `json:"password"`
	Roles     interface{} `json:"roles"`
	Address   string      `json:"address"`
	Buildings interface{} `json:"buildings"`
}
