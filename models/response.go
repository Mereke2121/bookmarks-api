package models

type ErrorResponse struct {
	Message string `json:"error"`
}

type ItemResponse struct {
	Status string `json:"status"`
}

type AddUserResponse struct {
	Id int `json:"id"`
}

type AuthorizationResponse struct {
	Token string `json:"token"`
}
