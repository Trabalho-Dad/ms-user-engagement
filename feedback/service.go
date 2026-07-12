package feedback

import (
	"bytes"
	"io"
	"net/http"
)

type Service struct {
	Repository Repository
}

func NewService(repo Repository) *Service {
	return &Service{Repository: repo}
}

func (s *Service) GetFeedbacksByFigureID(idFigure int) ([]Feedback, error) {
	return s.Repository.GetFeedbacksByFigureID(idFigure)
}

func (s *Service) CreateFeedback(feedback Feedback) (Feedback, error) {
	if err := feedback.Validate(); err != nil {
		return Feedback{}, err
	}

	url := "..."

	response, err := http.Get(url)
	if err != nil {
		return Feedback{}, ErrCouldNotValidateFigurePurchase
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return Feedback{}, ErrCouldNotValidateFigurePurchase
	}

	bodyString := string(bytes.TrimSpace(body))
	if bodyString == "[]" {
		return Feedback{}, ErrNoPurchaseFoundForFigure
	}

	return s.Repository.CreateFeedback(feedback)

}
