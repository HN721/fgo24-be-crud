package main

import (
	"minitask2/router"

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

	r.Run()

}
