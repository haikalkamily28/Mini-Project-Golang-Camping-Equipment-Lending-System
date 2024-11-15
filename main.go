package main

import (
	"log"
	"mini/config"
	loanRepository "mini/repository/loan"
    authRepository "mini/repository/auth"
    itemRepository "mini/repository/item"
	"mini/routes"
	"mini/service"

	"github.com/labstack/echo/v4"
)

func main() {
    config.ConnectDb()
    config.MigrateDB()
    e := echo.New()

    userRepo := authRepository.NewUserRepository(config.DB)
    userService := service.NewUserService(userRepo)

    loanRepo := loanRepository.NewLoanRepository(config.DB)
    loanService := service.NewLoanService(loanRepo)

    itemRepo := itemRepository.NewItemRepository(config.DB)
    itemService := service.NewItemService(itemRepo)

    routes.Routes(e, userService, loanService, itemService)

    log.Fatal(e.Start(":8080"))
}

