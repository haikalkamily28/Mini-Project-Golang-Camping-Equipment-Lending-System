package handler

import (
	"mini/entity"
	"mini/service"
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
    "github.com/golang-jwt/jwt/v5"
)

type LoanHandler struct {
    loanService *service.LoanService
}

func NewLoanHandler(loanService *service.LoanService) *LoanHandler {
    return &LoanHandler{loanService: loanService}
}


func (h *LoanHandler) GetAllLoans(c echo.Context) error {
    loans, err := h.loanService.GetAllLoans()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error fetching loans"})
    }
    return c.JSON(http.StatusOK, loans) // Return all loans with preloaded User and Item
}

func (h *LoanHandler) GetLoanByID(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid loan ID"})
    }

    loan, err := h.loanService.GetLoanByID(uint(id))
    if err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"message": "Loan not found"})
    }
    return c.JSON(http.StatusOK, loan) // Return loan with preloaded User and Item
}

func (h *LoanHandler) CreateLoan(c echo.Context) error {
    var loan entity.Loan
    if err := c.Bind(&loan); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
    }

    // Set the user ID from JWT claims
    user := c.Get("user").(*jwt.Token)
    claims := user.Claims.(jwt.MapClaims)
    userID := uint(claims["user_id"].(float64))
    loan.UserID = userID

    // Create the loan using the LoanService
    err := h.loanService.CreateLoan(&loan)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
    }

    // After creation, fetch the loan with preloaded User and Item
    loan, err = h.loanService.GetLoanByID(loan.ID)  // You should call GetLoanByID in LoanService
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error loading loan with User and Item"})
    }

    return c.JSON(http.StatusCreated, loan) // Return the loan with preloaded User and Item
}

func (h *LoanHandler) UpdateLoan(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid loan ID"})
    }

    user := c.Get("user").(*jwt.Token)
    claims := user.Claims.(jwt.MapClaims)
    userID := uint(claims["user_id"].(float64))

    loan, err := h.loanService.GetLoanByID(uint(id))
    if err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "Loan not found"})
    }

    if loan.UserID != userID {
        return c.JSON(http.StatusForbidden, map[string]string{"error": "You do not have permission to update this loan"})
    }

    if err := c.Bind(&loan); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
    }

    if loan.BorrowDate == "" || loan.ReturnDate == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "All fields are required"})
    }

    if err := h.loanService.UpdateLoan(&loan); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update loan"})
    }

    return c.JSON(http.StatusOK, loan)
}

func (h *LoanHandler) DeleteLoan(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid loan ID"})
    }

    user := c.Get("user").(*jwt.Token)
    claims := user.Claims.(jwt.MapClaims)
    userID := uint(claims["user_id"].(float64))

    loan, err := h.loanService.GetLoanByID(uint(id))
    if err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "Loan not found"})
    }

    if loan.UserID != userID {
        return c.JSON(http.StatusForbidden, map[string]string{"error": "You do not have permission to delete this loan"})
    }

    if err := h.loanService.DeleteLoan(uint(id)); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete loan"})
    }

    return c.JSON(http.StatusOK, map[string]string{"message": "Loan deleted successfully"})
}



