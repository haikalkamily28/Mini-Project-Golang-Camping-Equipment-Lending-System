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
	// Set a timeout context to prevent long queries
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result := r.db.WithContext(ctx).Preload("User").Preload("Item").Find(&loans)
	if result.Error != nil {
		log.Printf("Error fetching all loans: %v", result.Error)
		return nil, result.Error
	}
	return loans, nil
}

func (r *loanRepository) GetLoanByID(id uint) (entity.Loan, error) {
	var loan entity.Loan

	// Create a timeout context to prevent long queries
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Fetch loan with associated User and Item
	result := r.db.WithContext(ctx).Preload("User").Preload("Item").First(&loan, id)
	if result.Error != nil {
		log.Printf("Error preloading loan data for ID %d: %v", id, result.Error)
		// Return a more descriptive error
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return loan, errors.New("loan not found")
		}
		return loan, result.Error
	}

	log.Printf("Successfully fetched loan data for ID %d", id)
	return loan, nil
}

func (r *loanRepository) CreateLoan(loan *entity.Loan) error {
	// Create a timeout context for the operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Validate if the Item and User exist
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

	// Save loan to the database
	if err := r.db.WithContext(ctx).Create(loan).Error; err != nil {
		log.Printf("Failed to create loan: %v", err)
		return err
	}

	log.Printf("Loan created successfully with ID: %d", loan.ID)
	return nil
}

func (r *loanRepository) UpdateLoan(loan *entity.Loan) error {
	// Validate if the Item and User exist
	var item entity.Item
	if err := r.db.First(&item, loan.ItemID).Error; err != nil {
		log.Printf("Item with ID %d not found: %v", loan.ItemID, err)
		return errors.New("item not found")
	}

	var user entity.User
	if err := r.db.First(&user, loan.UserID).Error; err != nil {
		log.Printf("User with ID %d not found: %v", loan.UserID, err)
		return errors.New("user not found")
	}

	// Update loan in the database
	if err := r.db.Save(loan).Error; err != nil {
		log.Printf("Failed to update loan: %v", err)
		return err
	}

	log.Printf("Loan with ID %d updated successfully", loan.ID)
	return nil
}

func (r *loanRepository) DeleteLoan(id uint) error {
	var loan entity.Loan
	// Validate if the loan exists
	if err := r.db.First(&loan, id).Error; err != nil {
		log.Printf("Loan with ID %d not found: %v", id, err)
		return errors.New("loan not found")
	}

	// Delete loan from the database
	if err := r.db.Delete(&loan).Error; err != nil {
		log.Printf("Failed to delete loan: %v", err)
		return err
	}

	log.Printf("Loan with ID %d deleted successfully", id)
	return nil
}
