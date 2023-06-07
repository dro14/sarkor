package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Login    string `gorm:"unique"`
	Password string
	Name     string
	Age      int
}
