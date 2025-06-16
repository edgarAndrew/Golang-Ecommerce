package models

type Category struct {
	ID           uint   `gorm:"primaryKey"`
	CategoryName string
	CategoryIcon string
	SubCategories []SubCategory `gorm:"foreignKey:CategoryID"`
}
