package feedback_test

import (
	"reflect"
	"testing"
	"time"

	"ms-feedbacks/feedback"
	feedbackmocks "ms-feedbacks/feedback/mocks"
)

func TestService_GetFeedbacksByFigureID(t *testing.T) {
	t.Parallel()

	repo := feedbackmocks.NewRepository(t)
	expected := feedback.PaginatedFeedbacks{
		Feedbacks: []feedback.Feedback{
			{
				ID:          "1",
				Rating:      5,
				Description: "great",
				CreatedAt:   time.Date(2026, time.July, 12, 10, 0, 0, 0, time.UTC),
				IdFigure:    10,
				IdUser:      20,
			},
		},
		Page:       1,
		PageSize:   feedback.FeedbacksPageSize,
		TotalItems: 1,
		TotalPages: 1,
	}

	repo.EXPECT().GetFeedbacksByFigureID(10, 1).Return(expected, nil)

	service := feedback.NewService(repo)

	got, err := service.GetFeedbacksByFigureID(10, 1)
	if err != nil {
		t.Fatalf("GetFeedbacksByFigureID() error = %v", err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("GetFeedbacksByFigureID() = %#v, want %#v", got, expected)
	}
}

func TestService_GetFeedbackSummary(t *testing.T) {
	t.Parallel()

	repo := feedbackmocks.NewRepository(t)
	expected := feedback.Summary{
		TotalFeedbacks: 3,
		AverageRating:  4.5,
	}

	repo.EXPECT().GetFeedbackSummary(10).Return(expected, nil)

	service := feedback.NewService(repo)

	got, err := service.GetFeedbackSummary(10)
	if err != nil {
		t.Fatalf("GetFeedbackSummary() error = %v", err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("GetFeedbackSummary() = %#v, want %#v", got, expected)
	}
}

func TestService_CreateFeedback_ValidationError(t *testing.T) {
	t.Parallel()

	repo := feedbackmocks.NewRepository(t)
	service := feedback.NewService(repo)

	_, err := service.CreateFeedback(feedback.Feedback{})
	if err == nil {
		t.Fatal("CreateFeedback() expected validation error, got nil")
	}
}

func TestService_CreateFeedback_InvalidRating(t *testing.T) {
	t.Parallel()

	repo := feedbackmocks.NewRepository(t)
	service := feedback.NewService(repo)

	input := validFeedback()
	input.Rating = 6

	_, err := service.CreateFeedback(input)
	if err == nil {
		t.Fatal("CreateFeedback() expected validation error, got nil")
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

	service := feedback.NewService(repo)

	repo.EXPECT().CreateFeedback(input).Return(output, nil)

	got, err := service.CreateFeedback(input)
	if err != nil {
		t.Fatalf("CreateFeedback() error = %v", err)
	}

	if !reflect.DeepEqual(got, output) {
		t.Fatalf("CreateFeedback() = %#v, want %#v", got, output)
	}
}

func TestService_CreateFeedback_RepositoryError(t *testing.T) {
	t.Parallel()

	repo := feedbackmocks.NewRepository(t)
	input := validFeedback()

	service := feedback.NewService(repo)

	repo.EXPECT().CreateFeedback(input).Return(feedback.Feedback{}, feedback.ErrNoFeedbackFound)

	_, err := service.CreateFeedback(input)
	if err == nil {
		t.Fatal("CreateFeedback() expected error, got nil")
	}
}

func validFeedback() feedback.Feedback {
	return feedback.Feedback{
		Rating:      5,
		Description: "great product",
		IdFigure:    10,
		IdUser:      20,
	}
}
