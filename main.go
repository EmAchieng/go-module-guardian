package main

import (
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "module-guardian/handlers"
    "log"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found")
    }
	gin.SetMode(gin.ReleaseMode)

    r := gin.Default()
    r.GET("/users", handlers.GetUsers)
    r.Run(":8080")
}