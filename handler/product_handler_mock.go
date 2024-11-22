package handler

import (
	"SecretCare/entity"
	"github.com/stretchr/testify/mock"
)

// ProductMock is the mock implementation of the HandlerProduct interface
type ProductMock struct {
	mock.Mock
}

// CreateNewProduct mocks the CreateNewProduct method
func (m *ProductMock) CreateNewProduct(product entity.Product) error {
    args := m.Called(product) // Register the method call and retrieve mock arguments
    return args.Error(0)      // Return the error (or nil) specified in the mock setup
}

// GetProductsByTokoID mocks the GetProductsByTokoID method
func (m *ProductMock) GetProductsByTokoID(tokoID int) []entity.Product {
	args := m.Called(tokoID)
	return args.Get(0).([]entity.Product)
}

// DeleteProductById mocks the DeleteProductById method
func (m *ProductMock) DeleteProductById(id int) error {
    args := m.Called(id) // Get the mocked arguments
    return args.Error(0) // Return the first argument as an error
}

// UpdateStockById mocks the UpdateStockById method
func (m *ProductMock) UpdateStockById(id int, stock int) error {
    args := m.Called(id, stock) // Get the mocked arguments
    return args.Error(0)        // Return the first argument as an error
}

// GetProductReport mocks the GetProductReport method
func (m *ProductMock) GetProductReport(tokoID int) []entity.ProductReport {
	args := m.Called(tokoID)
	return args.Get(0).([]entity.ProductReport)
}

// GetProductsByTokoID mocks the GetProductsByTokoID method
func (m *ProductMock) GetAllProducts() ([]entity.Product, error) {
    args := m.Called()
    products, _ := args.Get(0).([]entity.Product)
    return products, args.Error(1)
}