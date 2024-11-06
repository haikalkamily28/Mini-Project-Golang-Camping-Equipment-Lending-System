package main

import (
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

    routes.RegisterRoutes(e, userService)

    e.Logger.Fatal(e.Start(":8080"))
}
