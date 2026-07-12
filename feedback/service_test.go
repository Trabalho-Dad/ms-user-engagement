package feedback_test

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"ms-feedbacks/feedback"
	feedbackmocks "ms-feedbacks/feedback/mocks"
)

func TestService_GetFeedbacksByFigureID(t *testing.T) {
	t.Parallel()

	repo := feedbackmocks.NewRepository(t)
	expected := []feedback.Feedback{
		{
			ID:          "1",
			CustomerID:  "customer-1",
			Rating:      5,
			Description: "great",
			CreatedAt:   time.Date(2026, time.July, 12, 10, 0, 0, 0, time.UTC),
			IdFigure:    10,
			IdUser:      20,
		},
	}

	repo.EXPECT().GetFeedbacksByFigureID(10).Return(expected, nil)

	service := feedback.NewService(repo)

	got, err := service.GetFeedbacksByFigureID(10)
	if err != nil {
		t.Fatalf("GetFeedbacksByFigureID() error = %v", err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("GetFeedbacksByFigureID() = %#v, want %#v", got, expected)
	}
}

func TestService_CreateFeedback_ValidationError(t *testing.T) {
	t.Parallel()

	repo := feedbackmocks.NewRepository(t)
	called := false
	service := feedback.NewServiceWithPurchaseCheck(repo, func() (bool, error) {
		called = true
		return true, nil
	})

	_, err := service.CreateFeedback(feedback.Feedback{})
	if err == nil {
		t.Fatal("CreateFeedback() expected validation error, got nil")
	}

	if called {
		t.Fatal("purchase check should not be called when validation fails")
	}
}

func TestService_CreateFeedback_NoPurchaseFound(t *testing.T) {
	t.Parallel()

	repo := feedbackmocks.NewRepository(t)
	service := feedback.NewServiceWithPurchaseCheck(repo, func() (bool, error) {
		return false, nil
	})

	_, err := service.CreateFeedback(validFeedback())
	if !errors.Is(err, feedback.ErrNoPurchaseFoundForFigure) {
		t.Fatalf("CreateFeedback() error = %v, want %v", err, feedback.ErrNoPurchaseFoundForFigure)
	}
}

func TestService_CreateFeedback_PurchaseValidationError(t *testing.T) {
	t.Parallel()

	repo := feedbackmocks.NewRepository(t)
	service := feedback.NewServiceWithPurchaseCheck(repo, func() (bool, error) {
		return false, errors.New("boom")
	})

	_, err := service.CreateFeedback(validFeedback())
	if !errors.Is(err, feedback.ErrCouldNotValidateFigurePurchase) {
		t.Fatalf("CreateFeedback() error = %v, want %v", err, feedback.ErrCouldNotValidateFigurePurchase)
	}
}

func TestService_CreateFeedback_Success(t *testing.T) {
	t.Parallel()

	repo := feedbackmocks.NewRepository(t)
	input := validFeedback()
	output := input
	output.ID = "feedback-1"
	output.CreatedAt = time.Date(2026, time.July, 12, 10, 30, 0, 0, time.UTC)
	output.UpdatedAt = time.Date(2026, time.July, 12, 10, 31, 0, 0, time.UTC)

	service := feedback.NewServiceWithPurchaseCheck(repo, func() (bool, error) {
		return true, nil
	})

	repo.EXPECT().CreateFeedback(input).Return(output, nil)

	got, err := service.CreateFeedback(input)
	if err != nil {
		t.Fatalf("CreateFeedback() error = %v", err)
	}

	if !reflect.DeepEqual(got, output) {
		t.Fatalf("CreateFeedback() = %#v, want %#v", got, output)
	}
}

func validFeedback() feedback.Feedback {
	return feedback.Feedback{
		ID:          "feedback-input-1",
		CustomerID:  "customer-1",
		Rating:      5,
		Description: "great product",
		CreatedAt:   time.Date(2026, time.July, 12, 10, 0, 0, 0, time.UTC),
		IdFigure:    10,
		IdUser:      20,
	}
}
