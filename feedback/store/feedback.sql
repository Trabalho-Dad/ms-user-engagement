-- name: GetFeedbacksByFigureID :many
SELECT 
    id, 
    description, 
    created_at, 
    updated_at, 
    id_figure, 
    id_user
FROM feedback
WHERE id_figure = $1;

-- name: CreateFeedback :one
INSERT INTO feedback (description, rating, id_figure, id_user)
VALUES ($1, $2, $3, $4)
RETURNING id, description, rating,created_at, updated_at, id_figure, id_user;