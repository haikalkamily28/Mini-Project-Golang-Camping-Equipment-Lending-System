package service

import (
    "mini/entity"
    "mini/repository"
)

type LoanService struct {
    loanRepo repository.LoanRepository
}

func NewLoanService(loanRepo repository.LoanRepository) *LoanService {
    return &LoanService{loanRepo: loanRepo}
}

func (s *LoanService) GetLoanRepository() repository.LoanRepository {
    return s.loanRepo
}

func (s *LoanService) GetAllLoans() ([]entity.Loan, error) {
    return s.loanRepo.GetAllLoans()
}

func (s *LoanService) GetLoanByID(id uint) (entity.Loan, error) {
    return s.loanRepo.GetLoanByID(id)
}

func (s *LoanService) CreateLoan(loan *entity.Loan) error {
    return s.loanRepo.CreateLoan(loan)
}

func (s *LoanService) UpdateLoan(loan *entity.Loan) error {
    return s.loanRepo.UpdateLoan(loan)
}

func (s *LoanService) DeleteLoan(id uint) error {
    return s.loanRepo.DeleteLoan(id)
}

