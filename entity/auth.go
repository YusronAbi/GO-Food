package entity

import (
	"context"
)

// AuthCredentials merepresentasikan data yang digunakan untuk login
type AuthCredentials struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// AuthRepository menangani operasi database terkait autentikasi
type AuthRepository interface {
	RegisterUser(ctx context.Context, user *User) (*User, error)
	GetUser(ctx context.Context, query interface{}, args ...interface{}) (*User, error)
}

// AuthService menangani logika bisnis autentikasi
type AuthService interface {
	Login(ctx context.Context, credentials *AuthCredentials) (token string, user *User, err error)
	Register(ctx context.Context, user *User) (token string, userData *User, err error)
	Logout(ctx context.Context, token string) error
}
