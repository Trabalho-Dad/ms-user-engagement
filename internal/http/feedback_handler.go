package http

import (
	"ms-feedbacks/feedback"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FeedbackHandler struct {
	FeedbackService feedback.UseCase
}

func NewFeedbackHandler(feedbackService feedback.UseCase) *FeedbackHandler {
	return &FeedbackHandler{
		FeedbackService: feedbackService,
	}
}

func (h *FeedbackHandler) GetFeedbacksByFigureID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idFigure, err := strconv.Atoi(c.Param("idFigure"))
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid figure id"})
			return
		}

		page := 1
		if p := c.Query("page"); p != "" {
			page, err = strconv.Atoi(p)
			if err != nil || page < 1 {
				c.JSON(400, gin.H{"error": "invalid page"})
				return
			}
		}

		feedbacks, err := h.FeedbackService.GetFeedbacksByFigureID(idFigure, page)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, feedbacks)
	}
}

func (h *FeedbackHandler) GetFeedbackSummary() gin.HandlerFunc {
	return func(c *gin.Context) {
		idFigure, err := strconv.Atoi(c.Param("idFigure"))
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid figure id"})
			return
		}

		summary, err := h.FeedbackService.GetFeedbackSummary(idFigure)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, summary)
	}
}

func (h *FeedbackHandler) CreateFeedback() gin.HandlerFunc {
	return func(c *gin.Context) {
		var feedback feedback.Feedback
		if err := c.ShouldBindJSON(&feedback); err != nil {
			c.JSON(400, gin.H{"error": "invalid request body"})
			return
		}

		createdFeedback, err := h.FeedbackService.CreateFeedback(feedback)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, createdFeedback)
	}
}
