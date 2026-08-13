package feedback

type Service struct {
	Repository Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		Repository: repo,
	}
}

func (s *Service) GetFeedbacksByFigureID(idFigure int, page int) (PaginatedFeedbacks, error) {
	return s.Repository.GetFeedbacksByFigureID(idFigure, page)
}

func (s *Service) GetFeedbackSummary(idFigure int) (Summary, error) {
	return s.Repository.GetFeedbackSummary(idFigure)
}

func (s *Service) CreateFeedback(feedback Feedback) (Feedback, error) {
	if err := feedback.Validate(); err != nil {
		return Feedback{}, err
	}
	return s.Repository.CreateFeedback(feedback)
}
