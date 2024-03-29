package models

// User struct
type User struct {
	Uid         string `json:"uid"`
	Email       string `json:"email"`
	Password    string `json:"password,omitempty"`
	PhoneNumber string `json:"phonenumber"`
	DisplayName string `json:"displayname"`
	PhotoURL    string `json:"photourl,omitempty"`
}
