package entity

type Loan struct {
    ID         uint   `json:"id" gorm:"primaryKey;autoIncrement"`
    ItemName   string `json:"item_name"`
    UserID     uint   `json:"user_id" gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
    User       User   `json:"user"`
    BorrowDate string `json:"borrow_date"`
    ReturnDate string `json:"return_date"`
    Status     string `json:"status"`
}
