package models

type Item struct {
	UserId int    `json:"-"`
	Id     int    `json:"id"`
	Url    string `json:"url"`
	Title  string `json:"title"`
}
