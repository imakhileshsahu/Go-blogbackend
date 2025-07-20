package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id        uint   `json:"id" gorm:"primaryKey"`
	FirstName string `json:"first_name" gorm:"not null"`
	LastName  string `json:"last_name" gorm:"not null"`
	Username  string `json:"username" gorm:"unique;not null"`
	Email     string `json:"email" gorm:"unique;not null"`
	Password  []byte `json:"-" gorm:"not null"`
	Phone     string `json:"phone" gorm:"not null"`
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashedPassword

}

func (user *User) CompparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}

// Password  string `json:"password" gorm:"not null"`
