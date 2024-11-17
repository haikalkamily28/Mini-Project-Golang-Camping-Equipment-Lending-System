package repository

import (
	"context"
	"errors"
	"log"
	"mini/entity"
	"time"
	"gorm.io/gorm"
)

type loanRepository struct {
	db *gorm.DB
}

func NewLoanRepository(db *gorm.DB) LoanRepository {
	return &loanRepository{db: db}
}

// In repository/loan_repository.go
func (r *loanRepository) GetAllLoans() ([]entity.Loan, error) {
    var loans []entity.Loan
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Ensure Preload is correctly used
    result := r.db.WithContext(ctx).Preload("User").Preload("Item").Find(&loans)
    if result.Error != nil {
        log.Printf("Error fetching all loans: %v", result.Error)
        return nil, result.Error
    }
    return loans, nil
}


func (r *loanRepository) GetLoanByID(id uint) (entity.Loan, error) {
	var loan entity.Loan
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result := r.db.WithContext(ctx).Preload("User").Preload("Item").First(&loan, id)
	if result.Error != nil {
		log.Printf("Error fetching loan by ID %d: %v", id, result.Error)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return loan, errors.New("loan not found")
		}
		return loan, result.Error
	}

	return loan, nil
}

func (r *loanRepository) CreateLoan(loan *entity.Loan) error {
    // Simpan loan terlebih dahulu
    result := r.db.Create(loan)
    if result.Error != nil {
        log.Printf("Error creating loan: %v", result.Error)
        return result.Error
    }

    // Preload relasi User dan Item setelah Loan dibuat
    if err := r.db.Preload("User").Preload("Item").First(loan, loan.ID).Error; err != nil {
        log.Printf("Error preloading related data: %v", err)
        return err
    }

    return nil
}

func (r *loanRepository) UpdateLoan(loan *entity.Loan) error {
	// Update loan in the database
	if err := r.db.Save(loan).Error; err != nil {
		return err
	}
	return nil
}

func (r *loanRepository) DeleteLoan(id uint) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result := r.db.WithContext(ctx).Delete(&entity.Loan{}, id)
	if result.Error != nil {
		log.Printf("Error deleting loan ID %d: %v", id, result.Error)
		return result.Error
	}

	return nil
}
