package models

type Item struct {
	Id    int    `json:"id"`
	Url   string `json:"url"`
	Title string `json:"title"`
}