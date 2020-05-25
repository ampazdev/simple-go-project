package entity

type User struct {
	ID          int    `json:"id"`
	Email       string `json:"email"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	JWT         JWT    `json:"jwt"`
}
