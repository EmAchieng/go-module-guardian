package handlers

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/gin-gonic/gin"
)

func TestGetUsers(t *testing.T) {
    gin.SetMode(gin.TestMode)
    r := gin.Default()
    r.GET("/users", GetUsers)

    req, _ := http.NewRequest("GET", "/users", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    if w.Code != http.StatusOK {
        t.Errorf("Expected status 200, got %d", w.Code)
    }
}