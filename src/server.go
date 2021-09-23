package main

import (
	"log"
	"net/http"

	"sample.com/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// router.StaticFS("/", http.Dir("./templates"))
	router.Static("/translate", "./templates")
	router.StaticFile("hogehoge", "./templates/hoge.html")
	// router.GET("/hello", hello, hello2)
	router.GET("/api/translate", controller.Translate)
	g := &gin.RouterGroup{Handlers: nil}
	log.Print(g)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server run failed.: ", err)
	}
}

func hello(c *gin.Context) {
	log.Print("Hello!")
	c.JSON(http.StatusOK, gin.H{"hoge": "Hello"})
}

func hello2(c *gin.Context) {
	log.Print("Hello2")
}
