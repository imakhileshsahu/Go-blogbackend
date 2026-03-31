// // package models

// // import "golang.org/x/crypto/bcrypt"

// <<<<<<< HEAD
// type User struct {
// 	Id        uint   `json:"id" gorm:"primaryKey"`
// 	FirstName string `json:"first_name" gorm:"not null"`
// 	LastName  string `json:"last_name" gorm:"not null"`
// 	// Username  string `json:"username" gorm:"unique;not null"`
// 	Email     string `json:"email" gorm:"unique;not null"`
// 	Password  []byte `json:"-" gorm:"not null"`
// 	Phone     string `json:"phone" gorm:"not null"`
// }
// =======
// // // type User struct {
// // // 	Id        uint    `json:"id" gorm:"primaryKey"`
// // // 	FirstName string  `json:"first_name" gorm:"not null"`
// // // 	LastName  string  `json:"last_name" gorm:"not null"`
// // // 	Username  *string `json:"username" gorm:"unique"`
// // // 	Email     string  `json:"email" gorm:"unique;not null"`
// // // 	Password  []byte  `json:"-" gorm:"not null"`
// // // 	Phone     string  `json:"phone" gorm:"not null"`
// // }
// >>>>>>> f45aa6b (migrated to postgres + auth fixe

// func (user *User) SetPassword(password string) {
// 	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
// 	user.Password = hashedPassword

// }

// func (user *User) CompparePassword(password string) error {
// 	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
// }

// Password  string `json:"password" gorm:"not null"`

package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	FirstName string  `json:"first_name" gorm:"not null"`
	LastName  string  `json:"last_name" gorm:"not null"`
	Username  *string `json:"username" gorm:"unique"`
	Email     string  `json:"email" gorm:"unique;not null"`
	Password  []byte  `json:"-" gorm:"not null"`
	Phone     string  `json:"phone" gorm:"not null"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (user *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return nil
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
