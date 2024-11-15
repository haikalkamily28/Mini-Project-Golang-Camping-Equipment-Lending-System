package service

import (
	"errors"
	"mini/entity"
	repository "mini/repository/item"
	"mini/utils"
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
	// Generate the description for the item using the Gemini API
	description, err := utils.CallGeminiAPI(item.Name)
	if err != nil {
		return errors.New("failed to generate description")
	}

	// Set the generated description to the item
	item.Description = description

	// Create the item in the repository
	return s.itemRepo.CreateItem(item)
}
