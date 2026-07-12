package main

import (
	"context"
	"fmt"
	"log"
	"ms-feedbacks/feedback"
	feedbackstore "ms-feedbacks/feedback/store"
	feedbackhttp "ms-feedbacks/internal/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	conn, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("Conectado ao PostgreSQL!")

	router := gin.Default()
	feedbackStore := feedbackstore.NewStore(conn)
	feedbackService := feedback.NewService(feedbackStore)
	feedbackHandler := feedbackhttp.NewFeedbackHandler(feedbackService)

	router.GET("/ms-feedback/get/:idFigure", feedbackHandler.GetFeedbacksByFigureID())

	router.POST("/ms-feedback", feedbackHandler.CreateFeedback())

	addr := os.Getenv("HTTP_ADDR")

	if err := router.Run(addr); err != nil {
		log.Fatal(err)
	}
}
