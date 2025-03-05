package models

type Wishlist struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"index"`
	ProductID uint   `gorm:"index"`
	Count     int    `gorm:"default:1"` // Added Count field with default value of 1
	User      User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // FK to User
}
