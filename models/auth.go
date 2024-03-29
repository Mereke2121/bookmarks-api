package models

type User struct {
	UserName string `json:"user_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Authorization struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
