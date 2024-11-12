package repository

import (
	"mini/entity"

	"gorm.io/gorm"
)

type LoanRepository interface {
    GetAllLoans() ([]entity.Loan, error)
    GetLoanByID(id uint) (entity.Loan, error)
    CreateLoan(loan *entity.Loan) error
    UpdateLoan(loan *entity.Loan) error
    DeleteLoan(id uint) error
}


type loanRepository struct {
    db *gorm.DB
}

func NewLoanRepository(db *gorm.DB) LoanRepository {
    return &loanRepository{db: db}
}

func (r *loanRepository) GetAllLoans() ([]entity.Loan, error) {
    var loans []entity.Loan
    result := r.db.Preload("User").Find(&loans)
    return loans, result.Error
}

func (r *loanRepository) GetLoanByID(id uint) (entity.Loan, error) {
    var loan entity.Loan
    result := r.db.Preload("User").First(&loan, id)
    return loan, result.Error
}

func (r *loanRepository) CreateLoan(loan *entity.Loan) error {
    return r.db.Create(loan).Error
}

func (r *loanRepository) UpdateLoan(loan *entity.Loan) error {
    return r.db.Save(loan).Error  
}

func (r *loanRepository) DeleteLoan(id uint) error {
    result := r.db.Delete(&entity.Loan{}, id)
    return result.Error
}
