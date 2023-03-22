package models

type ErrorResponse struct {
	Message string `json:"message"`
}

type RemoveItemResponse struct {
	Status string `json:"status"`
}

type AddItemResponse struct {
	Id int `json:"id"`
}

type AddUserResponse struct {
	Id int `json:"id"`
}

type AuthorizationResponse struct {
	Token string `json:"token"`
}
