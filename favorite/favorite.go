package favorite

import (
	"context"
	"errors"
)

const (
	ErrFavoriteAlreadyExists = "favorite already exists"
	ErrFavoriteNotFound      = "favorite not found"
)

var (
	ErrorFavoriteAlreadyExists = errors.New(ErrFavoriteAlreadyExists)
	ErrorFavoriteNotFound      = errors.New(ErrFavoriteNotFound)
)

type Repository interface {
	Create(ctx context.Context, favorite *Favorite) (*Favorite, error)
}

type UseCase interface {
	Create(ctx context.Context, favorite *Favorite) (*Favorite, error)
}
