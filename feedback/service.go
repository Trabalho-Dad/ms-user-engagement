package feedback

type Service struct {
	Repository Repository
}

func NewService(repo Repository) *Service {
	return &Service{Repository: repo}
}

func (s *Service) GetFeedbacksByFigureID(idFigure int) ([]Feedback, error) {
	return s.Repository.GetFeedbacksByFigureID(idFigure)
}
