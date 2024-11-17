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

func (s *ItemService) UpdateItem(id uint, item *entity.Item) error {
    // Cari item berdasarkan ID
    items, err := s.itemRepo.GetAllItems() // Menangani dua nilai yang dikembalikan
    if err != nil {
        return err // Jika ada error, return error
    }

    var existingItem entity.Item
    for _, i := range items {
        if i.ID == id {
            existingItem = i
            break
        }
    }

    if existingItem.ID == 0 {
        return errors.New("item not found") // Jika item tidak ditemukan
    }

    // Generate description using Gemini API
    description, err := utils.CallGeminiAPI(item.Name)
    if err != nil {
        return errors.New("failed to generate description for the item")
    }

    // Set the generated description to the item
    item.Description = description

    // Set ID for update
    item.ID = id
    return s.itemRepo.UpdateItem(item)
}


func (s *ItemService) DeleteItem(id uint) error {
    // Cek apakah item ada
    items, err := s.itemRepo.GetAllItems() // Menangani dua nilai yang dikembalikan
    if err != nil {
        return err // Jika ada error, return error
    }

    var existingItem entity.Item
    for _, i := range items {
        if i.ID == id {
            existingItem = i
            break
        }
    }

    if existingItem.ID == 0 {
        return errors.New("item not found") // Jika item tidak ditemukan
    }

    return s.itemRepo.DeleteItem(id)
}

