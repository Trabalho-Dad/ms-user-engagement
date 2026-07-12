package store

import (
	"context"
	"ms-feedbacks/feedback"
	"ms-feedbacks/internal/db"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
)

type Store struct {
	Queries Queries
}

type Queries interface {
	CreateFeedback(ctx context.Context, arg db.CreateFeedbackParams) (db.Feedback, error)
	GetFeedbacksByFigureID(ctx context.Context, idFigure pgtype.Int4) ([]db.Feedback, error)
}

func NewStore(database db.DBTX) *Store {
	return &Store{
		Queries: sqlcQueries{queries: db.New(database)},
	}
}

func (s *Store) GetFeedbacksByFigureID(idFigure int) ([]feedback.Feedback, error) {
	rows, err := s.Queries.GetFeedbacksByFigureID(context.Background(), pgtype.Int4{Int32: int32(idFigure), Valid: true})
	if err != nil {
		return nil, err
	}

	feedbacks := make([]feedback.Feedback, len(rows))
	for i, row := range rows {
		feedbacks[i] = toFeedback(row)
	}
	return feedbacks, nil
}

func (s *Store) CreateFeedback(fb feedback.Feedback) (feedback.Feedback, error) {
	arg := db.CreateFeedbackParams{
		Description: pgtype.Text{String: fb.Description, Valid: true},
		Rating:      int32(fb.Rating),
		IDFigure:    pgtype.Int4{Int32: int32(fb.IdFigure), Valid: true},
		IDUser:      pgtype.Int4{Int32: int32(fb.IdUser), Valid: true},
	}

	row, err := s.Queries.CreateFeedback(context.Background(), arg)
	if err != nil {
		return feedback.Feedback{}, err
	}

	return toFeedback(row), nil
}

type sqlcQueries struct {
	queries *db.Queries
}

func (q sqlcQueries) CreateFeedback(ctx context.Context, arg db.CreateFeedbackParams) (db.Feedback, error) {
	row, err := q.queries.CreateFeedback(ctx, arg)
	if err != nil {
		return db.Feedback{}, err
	}

	return db.Feedback{
		ID:          row.ID,
		Rating:      row.Rating,
		Description: row.Description,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
		IDFigure:    row.IDFigure,
		IDUser:      row.IDUser,
	}, nil
}

func (q sqlcQueries) GetFeedbacksByFigureID(ctx context.Context, idFigure pgtype.Int4) ([]db.Feedback, error) {
	rows, err := q.queries.GetFeedbacksByFigureID(ctx, idFigure)
	if err != nil {
		return nil, err
	}

	items := make([]db.Feedback, len(rows))
	for i, row := range rows {
		items[i] = db.Feedback{
			ID:          row.ID,
			Description: row.Description,
			CreatedAt:   row.CreatedAt,
			UpdatedAt:   row.UpdatedAt,
			IDFigure:    row.IDFigure,
			IDUser:      row.IDUser,
		}
	}

	return items, nil
}

func toFeedback(row db.Feedback) feedback.Feedback {
	return feedback.Feedback{
		ID:          strconv.FormatInt(int64(row.ID), 10),
		Description: row.Description.String,
		CreatedAt:   row.CreatedAt.Time,
		UpdatedAt:   row.UpdatedAt.Time,
		IdFigure:    int(row.IDFigure.Int32),
		IdUser:      int(row.IDUser.Int32),
	}
}
