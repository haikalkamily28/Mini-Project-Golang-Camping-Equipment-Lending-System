package routes

import (
	"mini/handler"
	authService "mini/service/auth"
	itemService "mini/service/item"
	loanService "mini/service/loan"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, userService authService.UserService, loanService *loanService.LoanService, itemService *itemService.ItemService) {
    userHandler := handler.UserHandler{UserService: userService}
    e.POST("/register", userHandler.Register)
    e.POST("/login", userHandler.Login)

    loanGroup := e.Group("/loans")
    loanGroup.Use(echojwt.WithConfig(echojwt.Config{
        SigningKey: []byte("aji"),
    }))
    loanHandler := handler.NewLoanHandler(loanService)
    loanGroup.GET("", loanHandler.GetAllLoans)
    loanGroup.GET("/:id", loanHandler.GetLoanByID)
    loanGroup.POST("", loanHandler.CreateLoan)
    loanGroup.PUT("/:id", loanHandler.UpdateLoan)
    loanGroup.DELETE("/:id", loanHandler.DeleteLoan)

    itemGroup := e.Group("/items")
    itemGroup.Use(echojwt.WithConfig(echojwt.Config{
        SigningKey: []byte("aji"),
    }))
    itemHandler := handler.NewItemHandler(itemService)
    itemGroup.GET("", itemHandler.GetAllItems)
    itemGroup.POST("", itemHandler.CreateItem)
}
