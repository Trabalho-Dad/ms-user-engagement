package http

import (
	"ms-feedbacks/favorite"

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
