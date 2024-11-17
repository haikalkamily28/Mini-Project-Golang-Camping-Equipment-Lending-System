package entity

type Loan struct {
    ID          uint       `json:"id"`
    UserID      uint       `json:"user_id"`
    ItemID      uint       `json:"item_id"`
    BorrowDate  string     `json:"borrow_date"`
    ReturnDate  string     `json:"return_date"`
    Status      string     `json:"status"`
    User        User       `gorm:"foreignKey:UserID"` // Ensure this relationship is correctly declared
    Item        Item       `gorm:"foreignKey:ItemID"` // Similarly for the Item
}


