package models

type Phone struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"index"`
	Phone       string `gorm:"unique" json:"phone" binding:"required,max=12"`
	Description string `json:"description" binding:"required"`
	IsMobile    bool   `json:"is_mobile" binding:"required"`
}
