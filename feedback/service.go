package feedback

import (
	"bytes"
	"io"
	"net/http"
)

type Service struct {
	Repository    Repository
	purchaseCheck func() (bool, error)
}

func NewService(repo Repository) *Service {
	return &Service{
		Repository:    repo,
		purchaseCheck: defaultPurchaseCheck,
	}
}

func NewServiceWithPurchaseCheck(repo Repository, purchaseCheck func() (bool, error)) *Service {
	if purchaseCheck == nil {
		purchaseCheck = defaultPurchaseCheck
	}

	return &Service{
		Repository:    repo,
		purchaseCheck: purchaseCheck,
	}
}

func (s *Service) GetFeedbacksByFigureID(idFigure int) ([]Feedback, error) {
	return s.Repository.GetFeedbacksByFigureID(idFigure)
}

func (s *Service) CreateFeedback(feedback Feedback) (Feedback, error) {
	if err := feedback.Validate(); err != nil {
		return Feedback{}, err
	}

	hasPurchase, err := s.purchaseCheck()
	if err != nil {
		return Feedback{}, ErrCouldNotValidateFigurePurchase
	}
	if !hasPurchase {
		return Feedback{}, ErrNoPurchaseFoundForFigure
	}

	return s.Repository.CreateFeedback(feedback)

}

func defaultPurchaseCheck() (bool, error) {
	url := "..."

	response, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return false, err
	}

	bodyString := string(bytes.TrimSpace(body))
	return bodyString != "[]", nil
}
