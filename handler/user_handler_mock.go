package handler

import (
	"SecretCare/entity"

	"github.com/stretchr/testify/mock"
)

type UserMock struct {
	mock.Mock
}

// GetUserByUsername mocks the GetUserByUsername method
func (m *UserMock) GetUserByUsername(username string) (*entity.Users, error) {
	args := m.Called(username)
	if user, ok := args.Get(0).(*entity.Users); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}

// UpdateMyAccount mocks the UpdateMyAccount method
func (m *UserMock) UpdateMyAccount(userId int, username, password, fullName *string) error {
	args := m.Called(userId, username, password, fullName)
	return args.Error(0)
}

// DeleteMyAccount mocks the DeleteMyAccount method
func (m *UserMock) DeleteMyAccount(userId int) error {
	args := m.Called(userId)
	return args.Error(0)
}

// ReportBuyerSpending mocks the ReportBuyerSpending method
func (m *UserMock) ReportBuyerSpending(userId int) ([]entity.UserBuyerReport, error) {
	args := m.Called(userId)
	if report, ok := args.Get(0).([]entity.UserBuyerReport); ok {
		return report, args.Error(1)
	}
	return nil, args.Error(1)
}

// ReportUserWithHighestSpending mocks the ReportUserWithHighestSpending method
func (m *UserMock) ReportUserWithHighestSpending(tokoId int) ([]entity.UserReportHighestSpending, error) {
	args := m.Called(tokoId)
	if report, ok := args.Get(0).([]entity.UserReportHighestSpending); ok {
		return report, args.Error(1)
	}
	return nil, args.Error(1)
}
