package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type LoginResponse struct {
	ID    int64  `json:"id"`
	Token string `json:"token"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

type Account struct {
	ID                int64     `json:"id"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	Email             string    `json:"email"`
	EncryptedPassword string    `json:"-"`
	CreatedAt         time.Time `json:"createdAt"`
}

func (a *Account) ValidPassword(pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(a.EncryptedPassword), []byte(pw)) == nil
}

func NewAccount(firstName, lastName, password, email string) (*Account, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &Account{
		FirstName:         firstName,
		LastName:          lastName,
		Email:             email,
		EncryptedPassword: string(encpw),
		CreatedAt:         time.Now().UTC(),
	}, nil
}
