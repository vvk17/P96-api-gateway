package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vvk17/P96-api-gateway/pkg/config"
)

func main() {

	log.Println("Config is loading...")
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	log.Println("Config loaded:", c)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "привет, это заглавная страница")
	})

	r.Run(":3001")
}
