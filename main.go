package main

import (
	"fmt"
	"minitask2/router"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title           CRUD Users
// @version         1.0
// @description     This is a sample server call server.
// @BasePath /

func main() {
	r := gin.Default()
	godotenv.Load()
	router.CombineRouter(r)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(fmt.Sprintf(":%s", port))

}
