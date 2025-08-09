package handlers

import (
    "net/http"
    "module-guardian/models"
    "github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
    users := []models.User{
        {ID: 1, Name: "Alice"},
        {ID: 2, Name: "Bob"},
    }
    c.JSON(http.StatusOK, users)
}