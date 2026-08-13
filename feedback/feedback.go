package feedback

import "errors"

const (
	MinRating         = 1
	MaxRating         = 5
	FeedbacksPageSize = 6
)

var (
	ErrNoFeedbackFound                = errors.New("no feedback found")
	ErrCouldNotValidateFigurePurchase = errors.New("could not validate figure purchase")
	ErrNoPurchaseFoundForFigure       = errors.New("no purchase found for this figure")
	ErrInvalidFeedback                = errors.New("invalid feedback")
	ErrFigureOrUserNotFound           = errors.New("figure or user not found")
)

type UseCase interface {
	GetFeedbacksByFigureID(idFigure int, page int) (PaginatedFeedbacks, error)
	CreateFeedback(feedback Feedback) (Feedback, error)
	GetFeedbackSummary(idFigure int) (Summary, error)
}

type Repository interface {
	GetFeedbacksByFigureID(idFigure int, page int) (PaginatedFeedbacks, error)
	CreateFeedback(feedback Feedback) (Feedback, error)
	GetFeedbackSummary(idFigure int) (Summary, error)
}
