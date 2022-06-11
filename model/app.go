package model

type App struct {
	Base
	Name string `gorm:"size:355" json:"name"`
}
