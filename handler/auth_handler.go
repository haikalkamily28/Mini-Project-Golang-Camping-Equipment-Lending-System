package handler

import (
	"mini/entity"
	handler "mini/handler/request"
	"mini/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
    UserService service.UserService 
}

func NewUserHandler(e *echo.Echo, userService service.UserService) {
    handler := &UserHandler{UserService: userService}
    e.POST("/register", handler.Register)
}

func (h *UserHandler) Register(c echo.Context) error {
    user := new(entity.User)
    if err := c.Bind(user); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
    }

    if user.Email == "" || user.Password == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "All fields are required"})
    }

    if err := h.UserService.Register(user); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }

    return c.JSON(http.StatusOK, map[string]string{"message": "Registration successful"})
}

func (h *UserHandler) Login(c echo.Context) error {
    req := new(handler.LoginRequest)
    if err := c.Bind(req); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
    }

    token, err := h.UserService.Login(req.Email, req.Password)
    if err != nil {
        return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, map[string]string{"token": token})
}

