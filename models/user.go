package models

// User struct
type User struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
	Status string `json:"status"`
	Id     int    `json:"id"`
}
