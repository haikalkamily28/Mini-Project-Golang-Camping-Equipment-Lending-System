package repository

import (
	"mini/entity"
    "errors"
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
    result := r.db.Preload("User").Preload("Item").Find(&loans)
    return loans, result.Error
}

func (r *loanRepository) GetLoanByID(id uint) (entity.Loan, error) {
    var loan entity.Loan
    result := r.db.Preload("User").Preload("Item").First(&loan, id) // Preload User and Item
    return loan, result.Error
}

func (r *loanRepository) CreateLoan(loan *entity.Loan) error {
    var item entity.Item
    if err := r.db.First(&item, loan.ItemID).Error; err != nil {
        return errors.New("item not found")
    }
    var user entity.User
    if err := r.db.First(&user, loan.UserID).Error; err != nil {
        return errors.New("user not found")
    }
    return r.db.Create(loan).Error
}

func (r *loanRepository) UpdateLoan(loan *entity.Loan) error {
    var item entity.Item
    if err := r.db.First(&item, loan.ItemID).Error; err != nil {
        return errors.New("item not found")
    }
    var user entity.User
    if err := r.db.First(&user, loan.UserID).Error; err != nil {
        return errors.New("user not found")
    }
    return r.db.Save(loan).Error
}

func (r *loanRepository) DeleteLoan(id uint) error {
    var loan entity.Loan
    if err := r.db.First(&loan, id).Error; err != nil {
        return errors.New("loan not found")
    }
    return r.db.Delete(&loan).Error
}

