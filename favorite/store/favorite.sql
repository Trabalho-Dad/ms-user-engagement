-- name: CreateFavorite :one
INSERT INTO favorite (
    id_user,
    id_figure
)
VALUES ($1, $2)
RETURNING id_user, id_figure, created_at;
