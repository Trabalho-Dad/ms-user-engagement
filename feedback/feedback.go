package feedback

const (
	MinRating = 1
	MaxRating = 5

	ErrNoFeedbackFound = "no feedback found"
)

type UseCase interface {
	GetFeedbacksByFigureID(idFigure int) ([]Feedback, error)
}

type Repository interface {
	GetFeedbacksByFigureID(idFigure int) ([]Feedback, error)
}
