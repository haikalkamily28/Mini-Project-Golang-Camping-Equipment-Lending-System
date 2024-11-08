package repository

import (
	"mini/entity"

	"gorm.io/gorm"
)

type LoanRepository interface {
    GetAllLoans() ([]entity.Loan, error)
    GetLoanByID(id uint) (entity.Loan, error)
}

type loanRepository struct {
    db *gorm.DB
}

func NewLoanRepository(db *gorm.DB) LoanRepository {
    return &loanRepository{db: db}
}

func (r *loanRepository) GetAllLoans() ([]entity.Loan, error) {
    var loans []entity.Loan
    result := r.db.Preload("User").Find(&loans) // Preload User data
    return loans, result.Error
}

func (r *loanRepository) GetLoanByID(id uint) (entity.Loan, error) {
    var loan entity.Loan
    result := r.db.Preload("User").First(&loan, id) // Preload User data
    return loan, result.Error
}
