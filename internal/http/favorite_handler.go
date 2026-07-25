package http

import (
	"ms-feedbacks/favorite"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FavoriteHandler struct {
	FavoriteService favorite.UseCase
}

func NewFavoriteHandler(favoriteService favorite.UseCase) *FavoriteHandler {
	return &FavoriteHandler{
		FavoriteService: favoriteService,
	}
}

func (h *FavoriteHandler) CreateFavorite() gin.HandlerFunc {
	return func(c *gin.Context) {
		var fav favorite.Favorite
		if err := c.ShouldBindJSON(&fav); err != nil {
			c.JSON(400, gin.H{"error": "invalid request body"})
			return
		}

		created, err := h.FavoriteService.Create(c.Request.Context(), &fav)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, created)
	}
}

func (h *FavoriteHandler) DeleteFavorite() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := strconv.ParseInt(c.Param("userID"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid user id"})
			return
		}

		figureID, err := strconv.ParseInt(c.Param("figureID"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid figure id"})
			return
		}

		err = h.FavoriteService.Delete(c.Request.Context(), userID, figureID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.Status(204)
	}
}

func (h *FavoriteHandler) ListFavoritesByUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := strconv.ParseInt(c.Param("userID"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid user id"})
			return
		}

		favorites, err := h.FavoriteService.ListByUser(c.Request.Context(), userID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, favorites)
	}
}
