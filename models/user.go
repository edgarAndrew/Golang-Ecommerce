package models

import (
	"fmt"
	"regexp"

	"gorm.io/gorm"
)

type User struct {
	ID        uint       `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Phone     string
	Email     string `gorm:"unique"`
	Username  string `gorm:"unique"`
	Password  string
	Admin   bool      `gorm:"default:false"`
	Addresses []Address `gorm:"foreignKey:UserID"`
	Reviews   []Review  `gorm:"foreignKey:UserID"`
	Wishlist  []Wishlist `gorm:"foreignKey:UserID"`
	Cart      []Cart    `gorm:"foreignKey:UserID"`
	Orders    []Order   `gorm:"foreignKey:UserID"`
}

// Custom validation function
func (u *User) ValidateEmail(db *gorm.DB) error {
	// Regular expression for validating email
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	valid := regexp.MustCompile(emailRegex).MatchString(u.Email)
	if !valid {
		fmt.Println("Invalid email format")
		return gorm.ErrInvalidData
	}
	return nil
}

func (u *User) ValidatePasswordLength(db *gorm.DB) error {
	if len(u.Password) < 8 {
		fmt.Println("Password must be at least 8 characters long")
		return gorm.ErrInvalidData
	}
	return nil
}

// getter
func (u *User) IsAdmin() bool {
	return u.Admin
}

// BeforeCreate is a GORM hook that runs before creating a user record
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// Run both validations
	if res := u.ValidateEmail(tx); res != nil {
		return res // error
	}
	if res := u.ValidatePasswordLength(tx); res != nil {
		return res // error
	}
	return nil
}
