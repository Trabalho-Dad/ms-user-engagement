-- name: GetFeedbacksByFigureID :many
SELECT
    id,
    rating,
    description,
    created_at,
    updated_at,
    id_figure,
    id_user
FROM feedback
WHERE id_figure = $1
ORDER BY id ASC
LIMIT $2
OFFSET $3;

-- name: CountFeedbacksByFigureID :one
SELECT COUNT(*)::bigint AS total_count
FROM feedback
WHERE id_figure = $1;

-- name: CreateFeedback :one
INSERT INTO feedback (description, rating, id_figure, id_user)
VALUES ($1, $2, $3, $4)
RETURNING id, description, rating,created_at, updated_at, id_figure, id_user;

-- name: GetFeedbackSummary :one
SELECT
    COUNT(*)::bigint AS total_feedbacks,
    COALESCE(AVG(rating), 0)::float8 AS average_rating
FROM feedback;