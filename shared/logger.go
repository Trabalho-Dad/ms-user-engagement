package shared

import (
	"log"
	"os"
)

func NewLogger() *log.Logger {
	return log.New(os.Stdout, "[ms-feedbacks] ", log.LstdFlags|log.LUTC)
}
