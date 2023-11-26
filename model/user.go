package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Uuid     string `json:"uuid" gorm:"uniqueIndex;not null"`
	Username string `json:"username" gorm:"uniqueIndex;not null"`
	Names    string `json:"names" gorm:"not null"`
	Email    string `json:"email" gorm:"uniqueIndex;not null"`
	Password string `json:"password" gorm:"not null"`
	gorm.Model
}

type UpdateUser struct {
	Names string `json:"names" validate:"required,min=3"`
}

type DeleteUser struct {
	Password string `json:"password" validate:"required,min=8"`
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(hash), err

}

func ComparePassword(password, hash string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hash))

	return err == nil

}
