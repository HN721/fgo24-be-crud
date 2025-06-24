package main

import (
	"minitask2/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	godotenv.Load()
	router.CombineRouter(r)

	r.Run()

}
