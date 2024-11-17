package entity

type UserResponse struct {
    ID    uint   `json:"id"`
    Email string `json:"email"`
}

type LoanResponse struct {
    ID         uint        `json:"id"`
    User       UserResponse `json:"user"`
    Item       Item        `json:"item"`
    BorrowDate string      `json:"borrow_date"`
    ReturnDate string      `json:"return_date"`
    Status     string      `json:"status"`
}
