package routes

import (
    "mini/handler"
    "mini/service"
    "github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, userService service.UserService) {
    userHandler := handler.UserHandler{UserService: userService}  
    e.POST("/register", userHandler.Register)
}
