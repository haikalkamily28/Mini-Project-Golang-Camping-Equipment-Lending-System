package routes

import (
	"mini/handler"
	authService "mini/service/auth"
	itemService "mini/service/item"
	loanService "mini/service/loan"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"log"
	"os"
)

func Routes(e *echo.Echo, userService authService.UserService, loanService *loanService.LoanService, itemService *itemService.ItemService) {
	jwtKey := os.Getenv("JWT_SECRET_KEY")
	if jwtKey == "" {
		log.Fatal("JWT_SECRET_KEY not set in environment variables")
	}

	userHandler := handler.UserHandler{UserService: userService}
	e.POST("/register", userHandler.Register)
	e.POST("/login", userHandler.Login)

	loanGroup := e.Group("/loans")
	loanGroup.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(jwtKey),
	}))
	loanHandler := handler.NewLoanHandler(loanService)
	loanGroup.GET("", loanHandler.GetAllLoans)
	loanGroup.GET("/:id", loanHandler.GetLoanByID)
	loanGroup.POST("", loanHandler.CreateLoan)
	loanGroup.PUT("/:id", loanHandler.UpdateLoan)
	loanGroup.DELETE("/:id", loanHandler.DeleteLoan)

	itemGroup := e.Group("/items")
	itemGroup.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(jwtKey),
	}))
	itemHandler := handler.NewItemHandler(itemService)
	itemGroup.GET("", itemHandler.GetAllItems)
	itemGroup.POST("", itemHandler.CreateItem)
}
