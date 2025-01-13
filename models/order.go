package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID        uint   `gorm:"primaryKey"`     // will be used as PK
	UserID    uint   `gorm:"index;not null"` // foreign key to User table
	ProductID uint   `gorm:"index;not null"` // foreign key to Product table
	Quantity  int    `gorm:"not null"`
	Status    string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	User    User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // FK to User
	Product Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`  // FK to Product
}

// Validations

var validStatuses = map[string]bool{
	"Pending":    true, // implies is a valid status
	"Processing": true,
	"Shipped":    true,
	"Delivered":  true,
}

// BeforeCreate hook to validate the status
func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	if !validStatuses[o.Status] {
		return fmt.Errorf("failed to create order: %s", o.Status)
	}
	o.CreatedAt = time.Now()
	return nil
}

// BeforeUpdate hook to validate the status
func (o *Order) BeforeUpdate(tx *gorm.DB) (err error) {
	if !validStatuses[o.Status] {
		return fmt.Errorf("failed to update order: %s", o.Status)
	}
	o.UpdatedAt = time.Now()
	return nil
}
