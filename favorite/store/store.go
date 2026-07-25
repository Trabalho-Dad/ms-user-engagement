package store

import (
	"context"
	"ms-feedbacks/favorite"
	"ms-feedbacks/internal/db"
)

type Store struct {
	Queries Queries
}

type Queries interface {
	CreateFavorite(ctx context.Context, arg db.CreateFavoriteParams) (db.Favorite, error)
	DeleteFavorite(ctx context.Context, arg db.DeleteFavoriteParams) error
}

func NewStore(database db.DBTX) *Store {
	return &Store{
		Queries: sqlcQueries{queries: db.New(database)},
	}
}

func (s *Store) Create(ctx context.Context, f *favorite.Favorite) (*favorite.Favorite, error) {
	arg := db.CreateFavoriteParams{
		IDUser:   int32(f.UserID),
		IDFigure: int32(f.FigureID),
	}

	row, err := s.Queries.CreateFavorite(ctx, arg)
	if err != nil {
		return nil, err
	}

	return toFavorite(row), nil
}

func (s *Store) Delete(ctx context.Context, userID, figureID int64) error {
	return s.Queries.DeleteFavorite(ctx, db.DeleteFavoriteParams{
		IDUser:   int32(userID),
		IDFigure: int32(figureID),
	})
}

type sqlcQueries struct {
	queries *db.Queries
}

func (q sqlcQueries) CreateFavorite(ctx context.Context, arg db.CreateFavoriteParams) (db.Favorite, error) {
	return q.queries.CreateFavorite(ctx, arg)
}

func (q sqlcQueries) DeleteFavorite(ctx context.Context, arg db.DeleteFavoriteParams) error {
	return q.queries.DeleteFavorite(ctx, arg)
}

func toFavorite(row db.Favorite) *favorite.Favorite {
	return &favorite.Favorite{
		UserID:    int64(row.IDUser),
		FigureID:  int64(row.IDFigure),
		CreatedAt: row.CreatedAt.Time,
	}
}
