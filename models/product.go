package models

type Product struct {
	ID               uint   `gorm:"primaryKey"`
	BrandID          uint   `gorm:"index"`
	SubCategoryID    uint   `gorm:"index"`
	ProductName      string
	ProductDescription string
	CompanyReview    string
	Price           float64
	Stock           int
	SalePercentage  float64
	Images          []Image `gorm:"foreignKey:ProductID"`
	Reviews         []Review       `gorm:"foreignKey:ProductID"`
}
