package main

import (
	"context"
	"fmt"
	"log"
	"ms-feedbacks/favorite"
	favoritestore "ms-feedbacks/favorite/store"
	"ms-feedbacks/feedback"
	feedbackstore "ms-feedbacks/feedback/store"
	"ms-feedbacks/internal/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Might be trying to deploy in production... Continuing from environment variables.")
	}

	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_NAME"),
	)

	conn, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("Conectado ao PostgreSQL!")

	router := gin.Default()
	router.Use(http.CORSMiddleware())

	feedbackStore := feedbackstore.NewStore(conn)
	feedbackService := feedback.NewService(feedbackStore)
	feedbackHandler := http.NewFeedbackHandler(feedbackService)

	favoriteStore := favoritestore.NewStore(conn)
	favoriteService := favorite.NewService(favoriteStore)
	favoriteHandler := http.NewFavoriteHandler(favoriteService)

	router.GET("/ms-feedback/get/:idFigure", feedbackHandler.GetFeedbacksByFigureID())
	router.GET("/ms-feedback/summary/:idFigure", feedbackHandler.GetFeedbackSummary())
	router.POST("/ms-feedback", feedbackHandler.CreateFeedback())

	router.POST("/ms-favorite", favoriteHandler.CreateFavorite())
	router.DELETE("/ms-favorite/:userID/:figureID", favoriteHandler.DeleteFavorite())
	router.GET("/ms-favorite/:userID", favoriteHandler.ListFavoritesByUser())

	addr := os.Getenv("HTTP_ADDR")

	if err := router.Run(addr); err != nil {
		log.Fatal(err)
	}
}
