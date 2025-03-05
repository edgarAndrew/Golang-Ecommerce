package models

type SubCategory struct {
	ID             uint   `gorm:"primaryKey"`
	CategoryID     uint   `gorm:"index"`
	SubCategoryName string
	Products       []Product `gorm:"foreignKey:SubCategoryID"`
	Product   Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // FK to Product
}
