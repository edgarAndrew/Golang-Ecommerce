package models

type Review struct {
	ID         uint   `gorm:"primaryKey"`
	ProductID  uint   `gorm:"index"`
	UserID     uint   `gorm:"index"`
	ReviewText string
	Stars      float64
	Product   Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // FK to Product
	User      User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
