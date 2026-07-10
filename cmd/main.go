package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
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

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	fmt.Println("Conectado ao PostgreSQL!")

	gin := gin.Default()
	addr := os.Getenv("HTTP_ADDR")

	if err := gin.Run(addr); err != nil {
		log.Fatal(err)
	}
}
