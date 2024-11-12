package routes

import (
    "mini/handler"
    "mini/service"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo-jwt/v4"
)

func Routes(e *echo.Echo, userService service.UserService, loanService *service.LoanService) {
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
}
