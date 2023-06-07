package models

// User model for storing user data
type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Login    string `json:"login,omitempty" gorm:"unique"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name" gorm:"index"`
	Age      int    `json:"age"`
}
