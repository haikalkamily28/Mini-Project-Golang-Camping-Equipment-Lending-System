package main

import (
    "log"
    "mini/config"
    "mini/repository"
    "mini/routes"
    "mini/service"
    "github.com/labstack/echo/v4"
)

func main() {
    config.ConnectDb()
    config.MigrateDB()
    e := echo.New()

    userRepo := repository.NewUserRepository(config.DB)
    userService := service.NewUserService(userRepo)

    loanRepo := repository.NewLoanRepository(config.DB)
    loanService := service.NewLoanService(loanRepo)

    routes.Routes(e, userService, loanService)

    log.Fatal(e.Start(":8080"))
}
