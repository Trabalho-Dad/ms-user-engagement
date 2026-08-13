package store

import (
	"context"
	"errors"
	"ms-feedbacks/feedback"
	"ms-feedbacks/internal/db"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

type fakeQueries struct {
	createFeedbackErr error
}

func (f fakeQueries) CreateFeedback(ctx context.Context, arg db.CreateFeedbackParams) (db.Feedback, error) {
	if f.createFeedbackErr != nil {
		return db.Feedback{}, f.createFeedbackErr
	}
	return db.Feedback{}, nil
}

func (f fakeQueries) GetFeedbacksByFigureID(ctx context.Context, arg db.GetFeedbacksByFigureIDParams) ([]db.GetFeedbacksByFigureIDRow, error) {
	return nil, nil
}

func (f fakeQueries) CountFeedbacksByFigureID(ctx context.Context, idFigure pgtype.Int4) (int64, error) {
	return 0, nil
}

func (f fakeQueries) GetFeedbackSummary(ctx context.Context, idFigure pgtype.Int4) (db.GetFeedbackSummaryRow, error) {
	return db.GetFeedbackSummaryRow{}, nil
}

func TestStore_CreateFeedback_ForeignKeyViolation(t *testing.T) {
	t.Parallel()

	s := &Store{
		Queries: fakeQueries{
			createFeedbackErr: &pgconn.PgError{Code: pgForeignKeyViolationCode, ConstraintName: "feedback_id_figure_fkey"},
		},
	}

	_, err := s.CreateFeedback(feedback.Feedback{
		Rating:      5,
		Description: "great product",
		IdFigure:    999,
		IdUser:      20,
	})

	if !errors.Is(err, feedback.ErrFigureOrUserNotFound) {
		t.Fatalf("CreateFeedback() error = %v, want %v", err, feedback.ErrFigureOrUserNotFound)
	}
}

func TestStore_CreateFeedback_OtherError(t *testing.T) {
	t.Parallel()

	dbErr := errors.New("boom")
	s := &Store{
		Queries: fakeQueries{createFeedbackErr: dbErr},
	}

	_, err := s.CreateFeedback(feedback.Feedback{
		Rating:      5,
		Description: "great product",
		IdFigure:    10,
		IdUser:      20,
	})

	if !errors.Is(err, dbErr) {
		t.Fatalf("CreateFeedback() error = %v, want %v", err, dbErr)
	}
}
