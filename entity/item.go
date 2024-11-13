package entity

type Item struct {
    ID    uint   `json:"id" gorm:"primaryKey"`
    Name  string `json:"name"`
    Price int    `json:"price"`
   	Description string `json:"description"`
}
