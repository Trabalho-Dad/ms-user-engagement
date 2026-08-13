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
	GetFeedbacksByFigureID(ctx context.Context, arg db.GetFeedbacksByFigureIDParams) ([]db.GetFeedbacksByFigureIDRow, error)
	CountFeedbacksByFigureID(ctx context.Context, idFigure pgtype.Int4) (int64, error)
	GetFeedbackSummary(ctx context.Context, idFigure pgtype.Int4) (db.GetFeedbackSummaryRow, error)
}

func NewStore(database db.DBTX) *Store {
	return &Store{
		Queries: sqlcQueries{queries: db.New(database)},
	}
}

func (s *Store) GetFeedbacksByFigureID(idFigure int, page int) (feedback.PaginatedFeedbacks, error) {
	ctx := context.Background()
	idFigureParam := pgtype.Int4{Int32: int32(idFigure), Valid: true}

	totalItems, err := s.Queries.CountFeedbacksByFigureID(ctx, idFigureParam)
	if err != nil {
		return feedback.PaginatedFeedbacks{}, err
	}

	rows, err := s.Queries.GetFeedbacksByFigureID(ctx, db.GetFeedbacksByFigureIDParams{
		IDFigure: idFigureParam,
		Limit:    feedback.FeedbacksPageSize,
		Offset:   int32((page - 1) * feedback.FeedbacksPageSize),
	})
	if err != nil {
		return feedback.PaginatedFeedbacks{}, err
	}

	feedbacks := make([]feedback.FeedbackResponseData, len(rows))
	for i, row := range rows {
		feedbacks[i] = toFeedbackResponse(row)
	}

	totalPages := 0
	if totalItems > 0 {
		totalPages = int((totalItems + int64(feedback.FeedbacksPageSize) - 1) / int64(feedback.FeedbacksPageSize))
	}

	return feedback.PaginatedFeedbacks{
		Feedbacks:  feedbacks,
		Page:       page,
		PageSize:   feedback.FeedbacksPageSize,
		TotalItems: totalItems,
		TotalPages: totalPages,
	}, nil
}

func (s *Store) GetFeedbackSummary(idFigure int) (feedback.Summary, error) {
	idFigureParam := pgtype.Int4{Int32: int32(idFigure), Valid: true}

	row, err := s.Queries.GetFeedbackSummary(context.Background(), idFigureParam)
	if err != nil {
		return feedback.Summary{}, err
	}

	return feedback.Summary{
		TotalFeedbacks: row.TotalFeedbacks,
		AverageRating:  row.AverageRating,
	}, nil
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

func (q sqlcQueries) GetFeedbackSummary(ctx context.Context, idFigure pgtype.Int4) (db.GetFeedbackSummaryRow, error) {
	return q.queries.GetFeedbackSummary(ctx, idFigure)
}

func (q sqlcQueries) CountFeedbacksByFigureID(ctx context.Context, idFigure pgtype.Int4) (int64, error) {
	return q.queries.CountFeedbacksByFigureID(ctx, idFigure)
}

func (q sqlcQueries) GetFeedbacksByFigureID(ctx context.Context, arg db.GetFeedbacksByFigureIDParams) ([]db.GetFeedbacksByFigureIDRow, error) {
	rows, err := q.queries.GetFeedbacksByFigureID(ctx, arg)
	if err != nil {
		return nil, err
	}

	items := make([]db.GetFeedbacksByFigureIDRow, len(rows))
	for i, row := range rows {
		items[i] = db.GetFeedbacksByFigureIDRow{
			ID:          row.ID,
			Description: row.Description,
			Rating:      row.Rating,
			CreatedAt:   row.CreatedAt,
			UpdatedAt:   row.UpdatedAt,
			IDFigure:    row.IDFigure,
			IDUser:      row.IDUser,
			UserName:    row.UserName,
		}
	}

	return items, nil
}

func toFeedback(row db.Feedback) feedback.Feedback {
	return feedback.Feedback{
		ID:          strconv.FormatInt(int64(row.ID), 10),
		Description: row.Description.String,
		Rating:      int(row.Rating),
		CreatedAt:   row.CreatedAt.Time,
		UpdatedAt:   row.UpdatedAt.Time,
		IdFigure:    int(row.IDFigure.Int32),
		IdUser:      int(row.IDUser.Int32),
	}
}

func toFeedbackResponse(row db.GetFeedbacksByFigureIDRow) feedback.FeedbackResponseData {
	return feedback.FeedbackResponseData{
		Feedback: feedback.Feedback{
			ID:          strconv.FormatInt(int64(row.ID), 10),
			Description: row.Description.String,
			Rating:      int(row.Rating),
			CreatedAt:   row.CreatedAt.Time,
			UpdatedAt:   row.UpdatedAt.Time,
			IdFigure:    int(row.IDFigure.Int32),
			IdUser:      int(row.IDUser.Int32),
		},
		UserName: row.UserName.String,
	}
}
