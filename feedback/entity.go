package feedback

import (
	"fmt"
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

type FeedbackResponseData struct {
	Feedback
	UserName string `json:"user_name"`
}

type PaginatedFeedbacks struct {
	Feedbacks  []FeedbackResponseData `json:"feedbacks"`
	Page       int                    `json:"page"`
	PageSize   int                    `json:"page_size"`
	TotalItems int64                  `json:"total_items"`
	TotalPages int                    `json:"total_pages"`
}

func (f Feedback) Validate() error {
	if f.Rating < 1 || f.Rating > 5 {
		return fmt.Errorf("%w: rating must be between 1 and 5", ErrInvalidFeedback)
	}

	if f.Description == "" {
		return fmt.Errorf("%w: description is required", ErrInvalidFeedback)
	}

	if f.IdFigure <= 0 {
		return fmt.Errorf("%w: id_figure is required", ErrInvalidFeedback)
	}

	if f.IdUser <= 0 {
		return fmt.Errorf("%w: id_user is required", ErrInvalidFeedback)
	}
	return nil
}
