package models

type Cart struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"index"`
	ProductID uint `gorm:"index"`
	Count     int
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // FK to User
}
