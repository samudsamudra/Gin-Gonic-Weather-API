package main

import (
    "weather-api/config"
    "weather-api/routes"
    "weather-api/middleware"

    "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
    config.ConnectDB()
    gin.SetMode(gin.ReleaseMode)

    r := gin.Default()
    r.POST("/register", routes.Register)
    r.POST("/login", routes.Login)
    r.GET("/weather", routes.GetWeather)

    auth := r.Group("/")
    auth.Use(middleware.AuthMiddleware())
    {
        auth.POST("/favorites", routes.AddFavoriteLocation)
        auth.GET("/favorites", routes.GetFavoriteLocations)
    }

    r.Run(":8080")
    r.SetTrustedProxies(nil)
}
