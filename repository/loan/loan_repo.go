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

func (r *loanRepository) GetAllLoans() ([]entity.Loan, error) {
	var loans []entity.Loan
	result := r.db.Preload("User").Preload("Item").Find(&loans)
	return loans, result.Error
}

func (r *loanRepository) GetLoanByID(id uint) (entity.Loan, error) {
	var loan entity.Loan

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result := r.db.WithContext(ctx).Preload("User").Preload("Item").First(&loan, id)

	if result.Error != nil {
		log.Printf("Error preloading loan data for ID %d: %v", id, result.Error)
		return loan, result.Error
	}

	log.Printf("Successfully fetched loan data for ID %d", id)
	return loan, nil
}

func (r *loanRepository) CreateLoan(loan *entity.Loan) error {
    // Membuat context dengan timeout
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Validasi apakah ItemID dan UserID valid
    var item entity.Item
    if err := r.db.WithContext(ctx).First(&item, loan.ItemID).Error; err != nil {
        log.Printf("Item with ID %d not found: %v", loan.ItemID, err)
        return errors.New("item not found")
    }

    var user entity.User
    if err := r.db.WithContext(ctx).First(&user, loan.UserID).Error; err != nil {
        log.Printf("User with ID %d not found: %v", loan.UserID, err)
        return errors.New("user not found")
    }

    // Menyimpan loan ke database
    if err := r.db.WithContext(ctx).Create(loan).Error; err != nil {
        log.Printf("Failed to create loan: %v", err)
        return err
    }

    log.Printf("Loan created successfully with ID: %d", loan.ID)
    return nil
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
