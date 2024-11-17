package handler

import (
	"mini/entity"
	itemService "mini/service/item"
	"net/http"
    "strconv"
	"github.com/labstack/echo/v4"
)

type ItemHandler struct {
    itemService *itemService.ItemService
}

func NewItemHandler(itemService *itemService.ItemService) *ItemHandler {
    return &ItemHandler{itemService: itemService}
}

func (h *ItemHandler) GetAllItems(c echo.Context) error {
    items, err := h.itemService.GetAllItems()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch items"})
    }
    return c.JSON(http.StatusOK, items)
}

func (h *ItemHandler) CreateItem(c echo.Context) error {
    item := new(entity.Item)
    if err := c.Bind(item); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
    }

    if item.Name == "" || item.Price <= 0 {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "All fields are required and must be valid"})
    }

    if err := h.itemService.CreateItem(item); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create item"})
    }

    return c.JSON(http.StatusCreated, item)
}

func (h *ItemHandler) UpdateItem(c echo.Context) error {
    id := c.Param("id")
    item := new(entity.Item)
    if err := c.Bind(item); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
    }

    // Konversi id ke uint
    itemID, err := strconv.Atoi(id)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
    }

    // Panggil service untuk update item
    if err := h.itemService.UpdateItem(uint(itemID), item); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update item"})
    }

    return c.JSON(http.StatusOK, item)
}

func (h *ItemHandler) DeleteItem(c echo.Context) error {
    id := c.Param("id")

    // Konversi id ke uint
    itemID, err := strconv.Atoi(id)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
    }

    // Panggil service untuk delete item
    if err := h.itemService.DeleteItem(uint(itemID)); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete item"})
    }

    return c.JSON(http.StatusOK, map[string]string{"message": "Item deleted successfully"})
}

