package models

// Phone model for storing phone data
type Phone struct {
	ID          uint   `json:"phone_id" gorm:"primaryKey"`
	UserID      uint   `json:"user_id" gorm:"unique"`
	Phone       string `json:"phone" binding:"required,max=12" gorm:"index"`
	Description string `json:"description"`
	IsFax       bool   `json:"is_fax"`
}
