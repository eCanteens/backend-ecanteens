package product

import (
	"strconv"

	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/customerror"
	"github.com/eCanteens/backend-ecanteens/src/helpers/pagination"
)

type Service interface {
	checkFeedback(userId uint, productId string) (*bool, error)
	addFeedback(body *feedbackScheme, userId uint, productId string) error
	removeFeedback(userId uint, productId string) error
	getFavorite(userId uint, query *paginationQS) (*pagination.Pagination[models.Product], error)
	addFavorite(userId uint, productId string) error
	removeFavorite(userId uint, productId string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) addFeedback(body *feedbackScheme, userId uint, productId string) error {
	id, err := strconv.ParseUint(productId, 10, 32)

	if err != nil {
		return customerror.New("Id produk tidak valid", 400)
	}

	feedbacks, err := s.repo.checkFeedback(userId, uint(id))

	if err != nil {
		return customerror.GormError(err, "Ulasan")
	}

	if len(*feedbacks) > 0 {
		if err := s.repo.updateFeedback(*(*feedbacks)[0].Id, body); err != nil {
			return customerror.GormError(err, "Ulasan")
		}

		return nil
	} else {
		feedback := &models.ProductFeedback{
			UserId:    userId,
			ProductId: uint(id),
			IsLike:    *body.IsLike,
		}

		if err := s.repo.createFeedback(feedback); err != nil {
			return customerror.GormError(err, "Ulasan")
		}

		return nil
	}
}

func (s *service) checkFeedback(userId uint, productId string) (*bool, error) {
	id, err := strconv.ParseUint(productId, 10, 32)

	if err != nil {
		return nil, customerror.New("Id produk tidak valid", 400)
	}

	feedbacks, err := s.repo.checkFeedback(userId, uint(id))
	if err != nil {
		return nil, customerror.GormError(err, "Produk")
	}

	if len(*feedbacks) > 0 {
		return &(*feedbacks)[0].IsLike, nil
	}

	return nil, nil
}

func (s *service) removeFeedback(userId uint, productId string) error {
	id, err := strconv.ParseUint(productId, 10, 32)

	if err != nil {
		return customerror.New("Id produk tidak valid", 400)
	}

	if err := s.repo.deleteFeedback(userId, uint(id)); err != nil {
		return customerror.GormError(err, "Ulasan")
	}

	return nil
}

func (s *service) getFavorite(userId uint, query *paginationQS) (*pagination.Pagination[models.Product], error) {
	var result = pagination.New(models.Product{})

	if err := s.repo.findFavorite(result, userId, query); err != nil {
		return nil, customerror.GormError(err, "Produk")
	}

	return result, nil
}

func (s *service) addFavorite(userId uint, productId string) error {
	id, err := strconv.ParseUint(productId, 10, 32)

	if err != nil {
		return customerror.New("Id produk tidak valid", 400)
	}

	favorites := s.repo.checkFavorite(userId, uint(id))

	if len(*favorites) > 0 {
		return customerror.New("Produk sudah di dalam list favorit anda", 400)
	}

	favorite := &models.FavoriteProduct{
		UserId:    userId,
		ProductId: uint(id),
	}

	if err := s.repo.createFavorite(favorite); err != nil {
		return customerror.GormError(err, "Produk")
	}

	return nil
}

func (s *service) removeFavorite(userId uint, productId string) error {
	id, err := strconv.ParseUint(productId, 10, 32)

	if err != nil {
		return customerror.New("Id produk tidak valid", 400)
	}

	favorites := s.repo.checkFavorite(userId, uint(id))

	if len(*favorites) == 0 {
		return customerror.New("Produk tidak ada di dalam list favorit anda", 400)
	}

	if err := s.repo.deleteFavorite(userId, uint(id)); err != nil {
		return customerror.GormError(err, "Produk")
	}

	return nil
}
