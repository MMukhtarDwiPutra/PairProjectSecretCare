package handler

import (
	"SecretCare/entity"
	"context"
	"github.com/stretchr/testify/mock"
)

// HandlerAuthMock adalah mock untuk interface HandlerAuth
type AuthMock struct {
	mock.Mock
}

// RegisterUser adalah mock untuk fungsi RegisterUser
func (m *AuthMock) RegisterUser(ctx context.Context, user entity.Users) {
	m.Called(ctx, user)
}

// Mock untuk Login
func (m *AuthMock) Login(username, password string) (bool, string, context.Context, error) {
    // Memanggil metode mock dengan parameter yang diberikan
    args := m.Called(username, password)

    // Mengembalikan hasil sesuai dengan ekspektasi yang ditentukan sebelumnya
    return args.Bool(0), args.String(1), args.Get(2).(context.Context), args.Error(3)
}