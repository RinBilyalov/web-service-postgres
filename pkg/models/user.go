package models

type User struct {
	ID      uint    `json:"id" gorm:"primary_key"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Comment Comment `json:"comment"`
}

type Comment struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	UserId uint   `json:"user_id"`
	Text   string `json:"text"`
}
