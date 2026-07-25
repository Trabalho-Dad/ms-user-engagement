package favorite

import "time"

type Favorite struct {
	UserID    int64     `json:"user_id"`
	FigureID  int64     `json:"figure_id"`
	CreatedAt time.Time `json:"created_at"`
}

func NewFavorite(userID, figureID int64) *Favorite {
	return &Favorite{
		UserID:   userID,
		FigureID: figureID,
	}
}
