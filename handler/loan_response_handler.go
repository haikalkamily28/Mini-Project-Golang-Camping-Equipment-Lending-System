package handler

import "mini/entity"

type LoanResponse struct {
	ID         uint     `json:"id"`
	ItemName   string   `json:"item_name"`
	BorrowDate string   `json:"borrow_date"`
	ReturnDate string   `json:"return_date"`
	Status     string   `json:"status"`
	UserID		uint      `json:"user_id"`
}

type LoanUser struct {
	ID uint `json:"id"`
}

func ToLoanResponse(loan entity.Loan) LoanResponse {
	return LoanResponse{
		ID:         loan.ID,
		BorrowDate: loan.BorrowDate,
		ReturnDate: loan.ReturnDate,
		Status:     loan.Status,
		UserID: 	loan.UserID,
	}
}

func ToLoanResponseList(loans []entity.Loan) []LoanResponse {
	var loanResponses []LoanResponse
	for _, loan := range loans {
		loanResponses = append(loanResponses, ToLoanResponse(loan))
	}
	return loanResponses
}