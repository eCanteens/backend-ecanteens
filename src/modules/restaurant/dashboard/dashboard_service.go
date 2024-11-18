package dashboard

import "github.com/eCanteens/backend-ecanteens/src/helpers/customerror"

type Service interface {
	toggleOpen(restoId uint, body *toggleOpenScheme) error
	dashboard(restoId uint) (*dashboardDto, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) toggleOpen(restoId uint, body *toggleOpenScheme) error {
	if err := s.repo.toggleOpen(restoId, body.IsOpen); err != nil {
		return customerror.GormError(err, "Restoran")
	}

	return nil
}

func (s *service) dashboard(restoId uint) (*dashboardDto, error) {
	var dto dashboardDto

	if err := s.repo.summary(&dto.Summary, restoId); err != nil {
		return nil, customerror.GormError(err, "Restoran")
	}

	if err := s.repo.latestHistory(&dto.History, restoId); err != nil {
		return nil, customerror.GormError(err, "Restoran")
	}

	return &dto, nil
}