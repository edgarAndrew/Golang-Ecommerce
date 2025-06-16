package models

type Image struct {
	ID        uint   `gorm:"primaryKey"`
	Url       string
	Name	  string
	ProductID uint    `gorm:"index;not null"`                                // foreign key to Product table
	Product   *Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // FK to Product
}
