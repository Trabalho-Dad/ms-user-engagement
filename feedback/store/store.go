package store

import (
	"context"
	"database/sql"
	"ms-feedbacks/feedback"
	"ms-feedbacks/internal/db"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
)

type Store struct {
	Queries *db.Queries
	DB      *sql.DB
}

func NewStore(database db.DBTX) *Store {
	return &Store{
		Queries: db.New(database),
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

func toFeedback(row db.GetFeedbacksByFigureIDRow) feedback.Feedback {
	return feedback.Feedback{
		ID:          strconv.FormatInt(int64(row.ID), 10),
		Description: row.Description.String,
		CreatedAt:   row.CreatedAt.Time,
		UpdatedAt:   row.UpdatedAt.Time,
		IdFigure:    int(row.IDFigure.Int32),
		IdUser:      int(row.IDUser.Int32),
	}
}
