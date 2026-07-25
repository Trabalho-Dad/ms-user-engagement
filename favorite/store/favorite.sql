-- name: CreateFavorite :one
INSERT INTO favorite (
    id_user,
    id_figure
)
VALUES ($1, $2)
RETURNING id_user, id_figure, created_at;

-- name: DeleteFavorite :exec
DELETE FROM favorite
WHERE id_user = $1
  AND id_figure = $2;
