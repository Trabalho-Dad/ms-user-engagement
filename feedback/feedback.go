package feedback

import "errors"

const (
	MinRating = 1
	MaxRating = 5
)

var (
	ErrNoFeedbackFound                = errors.New("no feedback found")
	ErrCouldNotValidateFigurePurchase = errors.New("could not validate figure purchase")
	ErrNoPurchaseFoundForFigure       = errors.New("no purchase found for this figure")
)

type UseCase interface {
	GetFeedbacksByFigureID(idFigure int) ([]Feedback, error)
	CreateFeedback(feedback Feedback) (Feedback, error)
}

type Repository interface {
	GetFeedbacksByFigureID(idFigure int) ([]Feedback, error)
	CreateFeedback(feedback Feedback) (Feedback, error)
}
