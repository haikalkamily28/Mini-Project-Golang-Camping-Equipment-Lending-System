package handler

import (
	"mini/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type LoanHandler struct {
    loanService service.LoanService
}

func NewLoanHandler(loanService service.LoanService) *LoanHandler {
    return &LoanHandler{loanService: loanService}
}

func (h *LoanHandler) GetAllLoans(c echo.Context) error {
    loans, err := h.loanService.GetAllLoans()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error fetching loans"})
    }
    return c.JSON(http.StatusOK, loans)
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

    return c.JSON(http.StatusOK, loan)
}
