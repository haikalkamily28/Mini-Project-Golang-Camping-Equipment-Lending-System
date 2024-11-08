package routes

import (
    "mini/handler"
    "mini/service"
    "github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, userService service.UserService, loanService service.LoanService) {
    // User routes
    userHandler := handler.UserHandler{UserService: userService}
    e.POST("/register", userHandler.Register)
    e.POST("/login", userHandler.Login)

    // Loan routes
    loanHandler := handler.NewLoanHandler(loanService)
    e.GET("/loans", loanHandler.GetAllLoans)
    e.GET("/loans/:id", loanHandler.GetLoanByID)
}
