package feedback

import "context"

type Repository interface {
	Save(ctx context.Context, feedback Feedback) (Feedback, error)
	List(ctx context.Context) ([]Feedback, error)
}
