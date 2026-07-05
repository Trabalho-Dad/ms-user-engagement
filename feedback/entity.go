package feedback

import (
	"errors"
	"strings"
	"time"
)

type Feedback struct {
	ID         string    `json:"id"`
	CustomerID string    `json:"customer_id"`
	OrderID    string    `json:"order_id,omitempty"`
	Rating     int       `json:"rating"`
	Comment    string    `json:"comment"`
	CreatedAt  time.Time `json:"created_at"`
}

func New(id, customerID, orderID string, rating int, comment string, createdAt time.Time) (Feedback, error) {
	feedback := Feedback{
		ID:         strings.TrimSpace(id),
		CustomerID: strings.TrimSpace(customerID),
		OrderID:    strings.TrimSpace(orderID),
		Rating:     rating,
		Comment:    strings.TrimSpace(comment),
		CreatedAt:  createdAt,
	}

	if err := feedback.Validate(); err != nil {
		return Feedback{}, err
	}

	return feedback, nil
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

	if f.Comment == "" {
		return errors.New("comment is required")
	}

	if f.CreatedAt.IsZero() {
		return errors.New("created at is required")
	}

	return nil
}
