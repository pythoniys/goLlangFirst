package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	memoryStorage := NewMemoryStorage()
	handler := NewHandler(memoryStorage)

	r := gin.Default()

	r.POST("/employee", handler.CreateEmployee)
	r.GET("/employee/:id", handler.GetEmployee)
	r.PUT("/employee/:id", handler.UpdateEmployee)
	r.DELETE("/employee/:id", handler.DeleteEmployee)
	r.GET("/employee", handler.GetAllEmployees)

	r.Run()
}
