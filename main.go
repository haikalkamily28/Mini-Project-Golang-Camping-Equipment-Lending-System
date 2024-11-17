package main

import (
	"log"
	"mini/config"
	loanRepository "mini/repository/loan"
    authRepository "mini/repository/auth"
    itemRepository "mini/repository/item"
	"mini/routes"
    authService "mini/service/auth"
	itemService "mini/service/item"
	loanService "mini/service/loan"
	"github.com/labstack/echo/v4"
)

func main() {
    config.ConnectDb()
    config.MigrateDB()
    e := echo.New()

    userRepo := authRepository.NewUserRepository(config.DB)
    userService := authService.NewUserService(userRepo)

    loanRepo := loanRepository.NewLoanRepository(config.DB)
    loanService := loanService.NewLoanService(loanRepo)

    itemRepo := itemRepository.NewItemRepository(config.DB)
    itemService := itemService.NewItemService(itemRepo)

    routes.Routes(e, userService, loanService, itemService)

    log.Fatal(e.Start(":8000"))
}

