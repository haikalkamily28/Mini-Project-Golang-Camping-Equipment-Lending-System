package service

import (
	"mini/entity"
	"mini/repository"
)

type ItemService struct {
    itemRepo repository.ItemRepository
}

func NewItemService(itemRepo repository.ItemRepository) *ItemService {
    return &ItemService{itemRepo: itemRepo}
}

func (s *ItemService) GetAllItems() ([]entity.Item, error) {
    return s.itemRepo.GetAllItems()
}

func (s *ItemService) CreateItem(item *entity.Item) error {
    return s.itemRepo.CreateItem(item)
}
