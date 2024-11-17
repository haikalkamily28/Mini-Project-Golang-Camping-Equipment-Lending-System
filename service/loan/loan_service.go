package service

import (
	"mini/entity"
	repository "mini/repository/loan"
    "fmt"
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

func (s *LoanService) CreateLoan(loan *entity.Loan) (*entity.Loan, error) {
    err := s.loanRepo.CreateLoan(loan)
    if err != nil {
        return nil, err
    }

    return loan, nil
}

func (s *LoanService) UpdateLoan(id uint, loan *entity.Loan) (*entity.Loan, error) {
    // Fetch existing loan
    existingLoan, err := s.loanRepo.GetLoanByID(id)
    if err != nil {
        return nil, fmt.Errorf("loan not found")
    }

    // Update fields in the existing loan
    existingLoan.BorrowDate = loan.BorrowDate
    existingLoan.ReturnDate = loan.ReturnDate
    existingLoan.Status = loan.Status

    // Ensure user and item associations are correctly set
    if loan.User.ID != 0 {
        existingLoan.User = loan.User
    }
    if loan.Item.ID != 0 {
        existingLoan.Item = loan.Item
    }

    // Save the updated loan
    if err := s.loanRepo.UpdateLoan(&existingLoan); err != nil {
        return nil, fmt.Errorf("failed to update loan")
    }

    return &existingLoan, nil
}

func (s *LoanService) DeleteLoan(id uint) error {
    return s.loanRepo.DeleteLoan(id)
}
