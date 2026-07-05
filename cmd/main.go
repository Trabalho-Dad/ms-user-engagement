package main

import (
	"log"

	"ms-feedbacks/shared"

	"github.com/gin-gonic/gin"
)

func main() {
	service := gin.Default()
	addr := shared.GetString("HTTP_ADDR", ":5656")

	if err := service.Run(addr); err != nil {
		log.Fatal(err)
	}
}
