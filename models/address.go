package models

type Address struct {
	ID         uint   `gorm:"primaryKey"`
	UserID     int   `gorm:"index"`
	Street     string
	City       string
	State      string
	Country    string
	PostalCode string
	IsActive   bool `gorm:"default:true"`
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // FK to User
}
