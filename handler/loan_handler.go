package handler

import (
	"log"
	"mini/entity"
	loanService "mini/service/loan"
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
)

type LoanHandler struct {
    loanService *loanService.LoanService
}

func NewLoanHandler(loanService *loanService.LoanService) *LoanHandler {
    return &LoanHandler{loanService: loanService}
}

func (h *LoanHandler) GetAllLoans(c echo.Context) error {
	loans, err := h.loanService.GetAllLoans()
	if err != nil {
		log.Printf("Error fetching loans: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error fetching loans"})
	}

	// Convert loans to LoanResponse
	var loanResponses []entity.LoanResponse
	for _, loan := range loans {
		loanResponse := entity.LoanResponse{
			ID:         loan.ID,
			BorrowDate: loan.BorrowDate,
			ReturnDate: loan.ReturnDate,
			Status:     loan.Status,
			User: entity.UserResponse{
				ID:    loan.User.ID,
				Email: loan.User.Email,
			},
			Item: loan.Item,
		}
		loanResponses = append(loanResponses, loanResponse)
	}

	return c.JSON(http.StatusOK, loanResponses)
}

func (h *LoanHandler) GetLoanByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Invalid loan ID: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid loan ID"})
	}

	loan, err := h.loanService.GetLoanByID(uint(id))
	if err != nil {
		log.Printf("Loan not found for ID %d: %v", id, err)
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Loan not found"})
	}

	// Create response struct without password field
	loanResponse := entity.LoanResponse{
		ID:         loan.ID,
		BorrowDate: loan.BorrowDate,
		ReturnDate: loan.ReturnDate,
		Status:     loan.Status,
		User: entity.UserResponse{
			ID:    loan.User.ID,
			Email: loan.User.Email,
		},
		Item: loan.Item,
	}

	return c.JSON(http.StatusOK, loanResponse)
}

func (h *LoanHandler) CreateLoan(c echo.Context) error {
	var loan entity.Loan
	if err := c.Bind(&loan); err != nil {
		log.Printf("Failed to bind loan data: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid loan data"})
	}

	createdLoan, err := h.loanService.CreateLoan(&loan)
	if err != nil {
		log.Printf("Failed to create loan: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create loan"})
	}

	// Create response struct without password field
	loanResponse := entity.LoanResponse{
		ID:         createdLoan.ID,
		BorrowDate: createdLoan.BorrowDate,
		ReturnDate: createdLoan.ReturnDate,
		Status:     createdLoan.Status,
		User: entity.UserResponse{
			ID:    createdLoan.User.ID,
			Email: createdLoan.User.Email,
		},
		Item: createdLoan.Item,
	}

	return c.JSON(http.StatusOK, loanResponse)
}

func (h *LoanHandler) UpdateLoan(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid loan ID"})
    }

    var loan entity.Loan
    if err := c.Bind(&loan); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid loan data"})
    }

    // Call service to update loan
    updatedLoan, err := h.loanService.UpdateLoan(uint(id), &loan)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
    }

    // Create response struct without password field
    loanResponse := entity.LoanResponse{
        ID:         updatedLoan.ID,
        BorrowDate: updatedLoan.BorrowDate,
        ReturnDate: updatedLoan.ReturnDate,
        Status:     updatedLoan.Status,
        User: entity.UserResponse{
            ID:    updatedLoan.User.ID,
            Email: updatedLoan.User.Email,
        },
        Item: updatedLoan.Item,
    }

    return c.JSON(http.StatusOK, loanResponse)
}


func (h *LoanHandler) DeleteLoan(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Invalid loan ID: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid loan ID"})
	}

	err = h.loanService.DeleteLoan(uint(id))
	if err != nil {
		log.Printf("Error deleting loan: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error deleting loan"})
	}

	return c.NoContent(http.StatusNoContent)
}

