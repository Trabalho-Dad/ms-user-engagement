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