package main

import (
	"fmt"
	"log"

	"test_galileosky/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Dbname)
	fmt.Println(psqlInfo)

	r := gin.Default()
	api := r.Group("/api")
	{
		api.POST("/tasks", func(ctx *gin.Context) {})
		api.GET("/tasks", func(ctx *gin.Context) {})
		api.GET("/tasks/export", func(ctx *gin.Context) {})
	}

	log.Printf("Server started on port %s", cfg.Port)
	log.Fatal(r.Run(":8080"))
}
