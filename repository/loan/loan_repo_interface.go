package repository

import (
	"mini/entity"
)

type LoanRepository interface {
    GetAllLoans() ([]entity.Loan, error)
    GetLoanByID(id uint) (entity.Loan, error)
    CreateLoan(loan *entity.Loan) error
    UpdateLoan(loan *entity.Loan) error
    DeleteLoan(id uint) error
}

