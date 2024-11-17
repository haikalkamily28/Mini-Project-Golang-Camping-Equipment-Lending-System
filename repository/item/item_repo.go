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

func (r *itemRepository) UpdateItem(item *entity.Item) error {
    return r.db.Save(item).Error
}

func (r *itemRepository) DeleteItem(id uint) error {
    return r.db.Delete(&entity.Item{}, id).Error
}
