// models/product.go
package models

type Product struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string  `gorm:"type:text"`
	Price       float64 `gorm:"type:decimal(10,2); not null"`
	Orders      []Order `gorm:"foreignKey:ProductID"` // Reverse relationship
}
