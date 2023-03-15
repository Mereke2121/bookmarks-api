package models

type ItemResponse struct {
	Status string `json:"status"`
}

type AddUserResponse struct {
	Id int `json:"id"`
}

type AuthorizationResponse struct {
	Token string `json:"token"`
}
