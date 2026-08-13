package feedback

import (
	"errors"
	"time"
)

type Feedback struct {
	ID          string    `json:"id"`
	Rating      int       `json:"rating"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	IdFigure    int       `json:"id_figure,omitempty"`
	IdUser      int       `json:"id_user,omitempty"`
}

type Summary struct {
	TotalFeedbacks int64   `json:"total_feedbacks"`
	AverageRating  float64 `json:"average_rating"`
}

type PaginatedFeedbacks struct {
	Feedbacks  []Feedback `json:"feedbacks"`
	Page       int        `json:"page"`
	PageSize   int        `json:"page_size"`
	TotalItems int64      `json:"total_items"`
	TotalPages int        `json:"total_pages"`
}

func (f Feedback) Validate() error {
	if f.Rating < 1 || f.Rating > 5 {
		return errors.New("rating must be between 1 and 5")
	}

	if f.Description == "" {
		return errors.New("description is required")
	}
	return nil
}
