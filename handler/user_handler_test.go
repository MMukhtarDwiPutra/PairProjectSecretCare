package handler_test

import (
	"SecretCare/entity"
	"SecretCare/handler"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserByUsername(t *testing.T) {
	// Create a new mock instance of HandlerUser
	mockHandler := new(handler.UserMock)

	// Prepare expected values
	expectedUser := &entity.Users{
		ID:       1,
		Username: "john_doe",
		FullName: "John Doe",
		Role:     "Pembeli",
		Password: "hashedpassword",
		TokoID:   2,
	}

	// Set up the mock to expect the GetUserByUsername method to be called with "john_doe" and return the expected user
	mockHandler.On("GetUserByUsername", "john_doe").Return(expectedUser, nil)

	// Call the method on the mock
	result, err := mockHandler.GetUserByUsername("john_doe")

	// Assert that there is no error
	assert.NoError(t, err)

	// Assert that the returned result matches the expected value
	assert.Equal(t, expectedUser, result)

	// Assert that the mock expectations were met
	mockHandler.AssertExpectations(t)
}

func TestUpdateMyAccount(t *testing.T) {
	// Create a new mock instance of HandlerUser
	mockHandler := new(handler.UserMock)

	// Test data
	userId := 1
	username := "new_username"
	password := "new_password"
	fullName := "New Full Name"

	// Set up the mock to expect the UpdateMyAccount method to be called with the user data
	mockHandler.On("UpdateMyAccount", userId, &username, &password, &fullName).Return(nil)

	// Call the method on the mock
	err := mockHandler.UpdateMyAccount(userId, &username, &password, &fullName)

	// Assert that there is no error
	assert.NoError(t, err)

	// Assert that the mock expectations were met
	mockHandler.AssertExpectations(t)
}

func TestUpdateMyAccount_UpdateUsernameOnly(t *testing.T) {
	// Create a new mock instance of HandlerUser
	mockHandler := new(handler.UserMock)

	// Test data
	userId := 1
	newUsername := "new_username"
	var newPassword, newFullName *string

	// Set up the mock to expect the UpdateMyAccount method to be called with the new username
	mockHandler.On("UpdateMyAccount", userId, &newUsername, newPassword, newFullName).Return(nil)

	// Call the method on the mock
	err := mockHandler.UpdateMyAccount(userId, &newUsername, newPassword, newFullName)

	// Assert that there is no error
	assert.NoError(t, err)

	// Assert that the mock expectations were met
	mockHandler.AssertExpectations(t)
}

func TestUpdateMyAccount_UpdateUsernameAndPassword(t *testing.T) {
	// Create a new mock instance of HandlerUser
	mockHandler := new(handler.UserMock)

	// Test data
	userId := 1
	newUsername := "new_username"
	newPassword := "new_password"
	var newFullName *string

	// Set up the mock to expect the UpdateMyAccount method to be called with the new username and password
	mockHandler.On("UpdateMyAccount", userId, &newUsername, &newPassword, newFullName).Return(nil)

	// Call the method on the mock
	err := mockHandler.UpdateMyAccount(userId, &newUsername, &newPassword, newFullName)

	// Assert that there is no error
	assert.NoError(t, err)

	// Assert that the mock expectations were met
	mockHandler.AssertExpectations(t)
}

func TestUpdateMyAccount_UpdateAllFields(t *testing.T) {
	// Create a new mock instance of HandlerUser
	mockHandler := new(handler.UserMock)

	// Test data
	userId := 1
	newUsername := "new_username"
	newPassword := "new_password"
	newFullName := "New Full Name"

	// Set up the mock to expect the UpdateMyAccount method to be called with all fields updated
	mockHandler.On("UpdateMyAccount", userId, &newUsername, &newPassword, &newFullName).Return(nil)

	// Call the method on the mock
	err := mockHandler.UpdateMyAccount(userId, &newUsername, &newPassword, &newFullName)

	// Assert that there is no error
	assert.NoError(t, err)

	// Assert that the mock expectations were met
	mockHandler.AssertExpectations(t)
}

func TestUpdateMyAccount_NoFieldsToUpdate(t *testing.T) {
	// Create a new mock instance of HandlerUser
	mockHandler := new(handler.UserMock)

	// Test data
	userId := 1
	var newUsername, newPassword, newFullName *string

	// Set up the mock to expect the UpdateMyAccount method to be called with no changes
	mockHandler.On("UpdateMyAccount", userId, newUsername, newPassword, newFullName).Return(nil)

	// Call the method on the mock
	err := mockHandler.UpdateMyAccount(userId, newUsername, newPassword, newFullName)

	// Assert that there is no error (no fields to update, so no changes should be made)
	assert.NoError(t, err)

	// Assert that the mock expectations were met
	mockHandler.AssertExpectations(t)
}

func TestDeleteMyAccount(t *testing.T) {
	// Create a new mock instance of HandlerUser
	mockHandler := new(handler.UserMock)

	// Test data
	userId := 1

	// Set up the mock to expect the DeleteMyAccount method to be called with the userId
	mockHandler.On("DeleteMyAccount", userId).Return(nil)

	// Call the method on the mock
	err := mockHandler.DeleteMyAccount(userId)

	// Assert that there is no error
	assert.NoError(t, err)

	// Assert that the mock expectations were met
	mockHandler.AssertExpectations(t)
}

func TestReportBuyerSpending(t *testing.T) {
	// Create a new mock instance of HandlerUser
	mockHandler := new(handler.UserMock)

	// Prepare expected values
	expectedReport := []entity.UserBuyerReport{
		{
			OrderID:       1,
			UserID:        1,
			FullName:      "John Doe",
			TotalSpending: 100.50,
			TotalQuantity: 2,
		},
	}

	// Set up the mock to expect the ReportBuyerSpending method to be called with userId 1 and return the expected report
	mockHandler.On("ReportBuyerSpending", 1).Return(expectedReport, nil)

	// Call the method on the mock
	result, err := mockHandler.ReportBuyerSpending(1)

	// Assert that there is no error
	assert.NoError(t, err)

	// Assert that the returned report matches the expected report
	assert.Equal(t, expectedReport, result)

	// Assert that the mock expectations were met
	mockHandler.AssertExpectations(t)
}

func TestReportUserWithHighestSpending(t *testing.T) {
	// Create a new mock instance of HandlerUser
	mockHandler := new(handler.UserMock)

	// Prepare expected values
	expectedReport := []entity.UserReportHighestSpending{
		{
			UserId:        1,
			FullName:      "John Doe",
			TotalSpending: 500.00,
		},
	}

	// Set up the mock to expect the ReportUserWithHighestSpending method to be called with tokoId 2 and return the expected report
	mockHandler.On("ReportUserWithHighestSpending", 2).Return(expectedReport, nil)

	// Call the method on the mock
	result, err := mockHandler.ReportUserWithHighestSpending(2)

	// Assert that there is no error
	assert.NoError(t, err)

	// Assert that the returned report matches the expected report
	assert.Equal(t, expectedReport, result)

	// Assert that the mock expectations were met
	mockHandler.AssertExpectations(t)
}

func TestUpdateMyAccount_UpdateUsernameOnly_Fail(t *testing.T) {
	// Create a new mock instance of HandlerUser
	mockHandler := new(handler.UserMock)

	// Test data
	userId := 1
	newUsername := "new_username"
	var newPassword, newFullName *string

	// Set up the mock to expect the UpdateMyAccount method to be called with the new username, but simulate an error
	mockHandler.On("UpdateMyAccount", userId, &newUsername, newPassword, newFullName).Return(fmt.Errorf("database error"))

	// Call the method on the mock
	err := mockHandler.UpdateMyAccount(userId, &newUsername, newPassword, newFullName)

	// Assert that an error occurred
	assert.Error(t, err)

	// Assert that the mock expectations were met
	mockHandler.AssertExpectations(t)
}

func TestUpdateMyAccount_UpdateUsernameAndPassword_Fail(t *testing.T) {
	// Create a new mock instance of HandlerUser
	mockHandler := new(handler.UserMock)

	// Test data
	userId := 1
	newUsername := "new_username"
	newPassword := "new_password"
	var newFullName *string

	// Set up the mock to expect the UpdateMyAccount method to be called with the new username and password, but simulate an error
	mockHandler.On("UpdateMyAccount", userId, &newUsername, &newPassword, newFullName).Return(fmt.Errorf("database error"))

	// Call the method on the mock
	err := mockHandler.UpdateMyAccount(userId, &newUsername, &newPassword, newFullName)

	// Assert that an error occurred
	assert.Error(t, err)

	// Assert that the mock expectations were met
	mockHandler.AssertExpectations(t)
}

func TestUpdateMyAccount_UpdateAllFields_Fail(t *testing.T) {
	// Create a new mock instance of HandlerUser
	mockHandler := new(handler.UserMock)

	// Test data
	userId := 1
	newUsername := "new_username"
	newPassword := "new_password"
	newFullName := "New Full Name"

	// Set up the mock to expect the UpdateMyAccount method to be called with all fields updated, but simulate an error
	mockHandler.On("UpdateMyAccount", userId, &newUsername, &newPassword, &newFullName).Return(fmt.Errorf("database error"))

	// Call the method on the mock
	err := mockHandler.UpdateMyAccount(userId, &newUsername, &newPassword, &newFullName)

	// Assert that an error occurred
	assert.Error(t, err)

	// Assert that the mock expectations were met
	mockHandler.AssertExpectations(t)
}
