package models

type Brand struct {
	ID        uint   `gorm:"primaryKey"`
	BrandName string
	BrandIcon string
	Products  []Product `gorm:"foreignKey:BrandID"`
}
