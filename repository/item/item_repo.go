package repository

import (
	"mini/entity"
	"gorm.io/gorm"
)

type itemRepository struct {
    db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
    return &itemRepository{db: db}
}

func (r *itemRepository) GetAllItems() ([]entity.Item, error) {
    var items []entity.Item
    result := r.db.Find(&items)
    return items, result.Error
}

func (r *itemRepository) CreateItem(item *entity.Item) error {
    return r.db.Create(item).Error
}
