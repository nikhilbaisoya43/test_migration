package model

type User struct {
	Base
	Name     string `gorm:"size:255" json:"name"`
	Email    string `gorm:"size:255" json:"email"`
	Password string `gorm:"size:255" json:"password"`
}
