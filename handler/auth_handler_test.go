package handler

import (
	"SecretCare/entity"
	"SecretCare/helpers"
	"context"
	"fmt"

	"testing"
	"github.com/stretchr/testify/assert"
)

// TestLogin_Success menguji proses login yang berhasil
func TestLogin_Success(t *testing.T) {
    // Inisialisasi mock untuk HandlerUser
    userMock := &UserMock{}
    handlerUser := HandlerUser(userMock) // pastikan authMock juga dipersiapkan dengan baik
    ctx := context.Background()

    // Data untuk pengujian
    username := "mmukhtar"
    password := "mmukhtar"
    hashedPassword, _ := helpers.HashPassword(password)
    user := entity.Users{
        ID:       2,
        Username: username,
        Password: hashedPassword,
        FullName: "Test User",
        TokoID:   1,
        Role:     "penjual",
    }

    // Tentukan perilaku mock untuk GetUserByUsername pada handlerUserMock
    userMock.On("GetUserByUsername", username).Return(user, nil).Once()
    _, _ = handlerUser.GetUserByUsername(username)

    // Buat mock untuk HandlerAuth dan tentukan perilaku Login
    authMock := &AuthMock{}
    authMock.On("Login", username, password).Return(true, "penjual", ctx, nil).Once()

    // Buat handlerAuth dengan menginjeksi handlerUserMock ke dalamnya
    handlerAuth := HandlerAuth(authMock) // pastikan authMock juga dipersiapkan dengan baik

    // Tes Login
    success, role, ctx, err := handlerAuth.Login(username, password)

    // Verifikasi hasil
    assert.True(t, success, "Login seharusnya berhasil")
    assert.Equal(t, "penjual", role, "Role harus sesuai")
    assert.NotNil(t, ctx, "Context tidak boleh nil")
    assert.Nil(t, err, "Error seharusnya nil")

    // Verifikasi ekspektasi mock
    userMock.AssertExpectations(t)
    authMock.AssertExpectations(t)
}

// TestLogin_Failure menguji proses login yang gagal
func TestLogin_Failure(t *testing.T) {
	// Inisialisasi mock untuk HandlerUser
	userMock := &UserMock{}
	handlerUser := HandlerUser(userMock) // pastikan authMock juga dipersiapkan dengan baik
	ctx := context.Background()

	// Data untuk pengujian
	username := "testuser"
	password := "wrongpassword"
	hashedPassword, _ := helpers.HashPassword("correctpassword")
	user := entity.Users{
		ID:       1,
		Username: username,
		Password: hashedPassword,
		FullName: "Test User",
		TokoID:   1,
		Role:     "",
	}

	// Mock perilaku GetUserByUsername
	userMock.On("GetUserByUsername", username).Return(user, nil)
    _, _ = handlerUser.GetUserByUsername(username)

	// Buat mock untuk HandlerAuth dan tentukan perilaku Login
    authMock := &AuthMock{}
    authMock.On("Login", username, password).Return(false, "", ctx, fmt.Errorf("Error database")).Once()

    // Buat handlerAuth dengan menginjeksi handlerUserMock ke dalamnya
    handlerAuth := HandlerAuth(authMock) // pastikan authMock juga dipersiapkan dengan baik

	// Tes Login
	success, role, ctx, err := handlerAuth.Login(username, password)

	// Verifikasi hasil
	assert.False(t, success, "Login seharusnya gagal")
	assert.Empty(t, role, "Role harus kosong")
	assert.NotNil(t, ctx, "Context tidak boleh nil")
	assert.NotNil(t, err, "Error seharusnya tidak nil")

	// Verifikasi ekspektasi mock
	userMock.AssertExpectations(t)
}
