package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/hello")

	if err := router.Run("localhost:8080"); err != nil {
		log.Fatal("Sever run failed.: ", err)
	}
}

func hello() {
	log.Print("Hello!")
}
