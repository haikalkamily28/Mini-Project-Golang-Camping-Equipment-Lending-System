package entity

type Loan struct {
    ID         uint   `json:"id" gorm:"primaryKey"`
    UserID     uint   `json:"user_id"`
    ItemID     uint   `json:"item_id"`
    BorrowDate string `json:"borrow_date"`
    ReturnDate string `json:"return_date"`
    Status     string `json:"status"` 

    User       User   `json:"user" gorm:"foreignKey:UserID"`
    Item       Item   `json:"item" gorm:"foreignKey:ItemID"`
}
