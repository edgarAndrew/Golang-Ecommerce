package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID         uint       `gorm:"primaryKey"`
	UserID     uint       `gorm:"index;not null"`
	TotalPrice float64    `gorm:"not null"`
	Status     string     `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time

	User       User        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID"`
}

// Valid order statuses
var validOrderStatuses = map[string]bool{
	"Pending":    true,
	"Processing": true,
	"Shipped":    true,
	"Delivered":  true,
	"Cancelled":  true,
}

// BeforeCreate hook to validate status
func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	if !validOrderStatuses[o.Status] {
		return fmt.Errorf("failed to create order: invalid status '%s'", o.Status)
	}
	o.CreatedAt = time.Now()
	return nil
}

// BeforeUpdate hook to validate status
func (o *Order) BeforeUpdate(tx *gorm.DB) (err error) {
	if !validOrderStatuses[o.Status] {
		return fmt.Errorf("failed to update order: invalid status '%s'", o.Status)
	}
	o.UpdatedAt = time.Now()
	return nil
}
