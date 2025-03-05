package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderItem struct {
	ID        uint    `gorm:"primaryKey"`
	OrderID   uint    `gorm:"index;not null"`
	ProductID uint    `gorm:"index;not null"`
	Quantity  int     `gorm:"not null"`
	Price     float64 `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Order   Order   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Product Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// BeforeCreate hook to set timestamps
func (oi *OrderItem) BeforeCreate(tx *gorm.DB) (err error) {
	oi.CreatedAt = time.Now()
	return nil
}

// BeforeUpdate hook to update timestamps
func (oi *OrderItem) BeforeUpdate(tx *gorm.DB) (err error) {
	oi.UpdatedAt = time.Now()
	return nil
}
