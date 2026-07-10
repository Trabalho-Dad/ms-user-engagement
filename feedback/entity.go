package feedback

import (
	"errors"
	"time"
)

type Feedback struct {
	ID          string    `json:"id"`
	CustomerID  string    `json:"customer_id"`
	Rating      int       `json:"rating"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	IdFigure    int       `json:"id_figure,omitempty"`
	IdUser      int       `json:"id_user,omitempty"`
}

func (f Feedback) Validate() error {
	if f.ID == "" {
		return errors.New("feedback id is required")
	}

	if f.CustomerID == "" {
		return errors.New("customer id is required")
	}

	if f.Rating < 1 || f.Rating > 5 {
		return errors.New("rating must be between 1 and 5")
	}

	if f.Description == "" {
		return errors.New("description is required")
	}

	if f.CreatedAt.IsZero() {
		return errors.New("created at is required")
	}
	return nil
}
