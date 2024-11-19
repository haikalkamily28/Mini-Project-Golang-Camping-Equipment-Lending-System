package service_test

import (
	"errors"
	"testing"

	"mini/entity"
	loanService "mini/service/loan"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockLoanRepository struct {
	mock.Mock
}

func (m *MockLoanRepository) GetAllLoans() ([]entity.Loan, error) {
	args := m.Called()
	return args.Get(0).([]entity.Loan), args.Error(1)
}

func (m *MockLoanRepository) GetLoanByID(id uint) (entity.Loan, error) {
	args := m.Called(id)
	return args.Get(0).(entity.Loan), args.Error(1)
}

func (m *MockLoanRepository) CreateLoan(loan *entity.Loan) error {
	args := m.Called(loan)
	return args.Error(0)
}

func (m *MockLoanRepository) UpdateLoan(loan *entity.Loan) error {
	args := m.Called(loan)
	return args.Error(0)
}

func (m *MockLoanRepository) DeleteLoan(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestLoanService_GetAllLoans(t *testing.T) {
	mockRepo := new(MockLoanRepository)
	mockLoans := []entity.Loan{
		{
			ID:         1,
			UserID:     1,
			ItemID:     1,
			BorrowDate: "2024-11-10",
			ReturnDate: "2024-11-20",
			Status:     "Borrowed",
			User:       entity.User{ID: 1, Email: "ajit@mail.com"},
			Item:       entity.Item{ID: 1, Name: "tongkat daki"},
		},
		{
			ID:         2,
			UserID:     2,
			ItemID:     2,
			BorrowDate: "2024-11-11",
			ReturnDate: "2024-11-21",
			Status:     "Returned",
			User:       entity.User{ID: 2, Email: "ajit@mail.com"},
			Item:       entity.Item{ID: 2, Name: "tongkat daki"},
		},
	}
	mockRepo.On("GetAllLoans").Return(mockLoans, nil)

	loanService := loanService.NewLoanService(mockRepo)
	loans, err := loanService.GetAllLoans()

	assert.NoError(t, err)
	assert.Equal(t, mockLoans, loans)
	mockRepo.AssertExpectations(t)
}

func TestLoanService_GetLoanByID(t *testing.T) {
	mockRepo := new(MockLoanRepository)
	mockLoan := entity.Loan{
		ID:         1,
		UserID:     1,
		ItemID:     1,
		BorrowDate: "2024-11-10",
		ReturnDate: "2024-11-20",
		Status:     "Borrowed",
		User:       entity.User{ID: 1, Email: "ajit@mail.com"},
		Item:       entity.Item{ID: 1, Name: "tongkat daki"},
	}
	mockRepo.On("GetLoanByID", uint(1)).Return(mockLoan, nil)

	loanService := loanService.NewLoanService(mockRepo)
	loan, err := loanService.GetLoanByID(1)

	assert.NoError(t, err)
	assert.Equal(t, mockLoan, loan)
	mockRepo.AssertExpectations(t)
}

func TestLoanService_CreateLoan(t *testing.T) {
	mockRepo := new(MockLoanRepository)
	mockLoan := &entity.Loan{
		ID:         1,
		UserID:     1,
		ItemID:     1,
		BorrowDate: "2024-11-10",
		ReturnDate: "2024-11-20",
		Status:     "Borrowed",
		User:       entity.User{ID: 1, Email: "ajit@mail.com"},
		Item:       entity.Item{ID: 1, Name: "tongkat daki"},
	}

	mockRepo.On("CreateLoan", mockLoan).Return(mockLoan, nil)

	loanService := loanService.NewLoanService(mockRepo)
	createdLoan, err := loanService.CreateLoan(mockLoan)

	assert.NoError(t, err)
	assert.Equal(t, mockLoan, createdLoan)
	mockRepo.AssertExpectations(t)
}


func TestLoanService_UpdateLoan(t *testing.T) {
	mockRepo := new(MockLoanRepository)
	mockLoan := &entity.Loan{
		ID:         1,
		UserID:     1,
		ItemID:     1,
		BorrowDate: "2024-11-10",
		ReturnDate: "2024-11-20",
		Status:     "Returned",
		User:       entity.User{ID: 1, Email: "ajit@mail.com"},
		Item:       entity.Item{ID: 1, Name: "tongkat daki"},
	}

	mockRepo.On("UpdateLoan", mockLoan).Return(mockLoan, nil)

	loanService := loanService.NewLoanService(mockRepo)
	updatedLoan, err := loanService.UpdateLoan(mockLoan.ID, mockLoan) 

	assert.NoError(t, err)
	assert.Equal(t, mockLoan, updatedLoan) 
	mockRepo.AssertExpectations(t)
}


func TestLoanService_DeleteLoan(t *testing.T) {
	mockRepo := new(MockLoanRepository)
	mockRepo.On("DeleteLoan", uint(1)).Return(nil)

	loanService := loanService.NewLoanService(mockRepo)
	err := loanService.DeleteLoan(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestLoanService_GetLoanByID_NotFound(t *testing.T) {
	mockRepo := new(MockLoanRepository)
	mockRepo.On("GetLoanByID", uint(1)).Return(entity.Loan{}, errors.New("loan not found"))

	loanService := loanService.NewLoanService(mockRepo)
	loan, err := loanService.GetLoanByID(1)

	assert.Error(t, err)
	assert.EqualError(t, err, "loan not found")
	assert.Equal(t, entity.Loan{}, loan)
	mockRepo.AssertExpectations(t)
}
