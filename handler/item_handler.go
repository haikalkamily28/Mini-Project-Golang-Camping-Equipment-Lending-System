package handler

import (
    "mini/entity"
    "mini/service"
    "net/http"
    "github.com/labstack/echo/v4"
)

type ItemHandler struct {
    itemService *service.ItemService
}

func NewItemHandler(itemService *service.ItemService) *ItemHandler {
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

    if item.Name == "" || item.Description == "" || item.Price <= 0 {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "All fields are required and must be valid"})
    }

    if err := h.itemService.CreateItem(item); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create item"})
    }

    return c.JSON(http.StatusCreated, item)
}
