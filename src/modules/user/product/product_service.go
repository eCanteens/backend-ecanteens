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

	feedback, err := s.repo.checkFeedback(userId, uint(id))

	if err != nil {
		return customerror.GormError(err, "Ulasan")
	}

	if body.IsLiked == nil {
		// Remove feedback
		if feedback == nil {
			return customerror.New("Produk belum dilike/didislike", 400)
		}

		if err := s.repo.deleteFeedback(userId, uint(id)); err != nil {
			return customerror.GormError(err, "Ulasan")
		}

		return nil
	}

	if feedback != nil {
		// Update feedback
		if feedback.IsLike == *body.IsLiked {
			msg := "Produk sudah di"
			if feedback.IsLike {
				msg += "like"
			} else {
				msg += "dislike"
			}

			return customerror.New(msg, 400)
		}

		if err := s.repo.updateFeedback(*feedback.Id, body); err != nil {
			return customerror.GormError(err, "Ulasan")
		}

		return nil
	} else {
		// Create feedback
		feedback := &models.ProductFeedback{
			UserId:    userId,
			ProductId: uint(id),
			IsLike:    *body.IsLiked,
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

	feedback, err := s.repo.checkFeedback(userId, uint(id))
	if err != nil {
		return nil, customerror.GormError(err, "Ulasan")
	}

	if feedback != nil {
		return &feedback.IsLike, nil
	}

	return nil, nil
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

	favorite, err := s.repo.checkFavorite(userId, uint(id))
	if err != nil {
		return customerror.GormError(err, "Produk")
	}

	if favorite != nil {
		return customerror.New("Produk sudah di dalam list favorit anda", 400)
	}

	newFavorite := &models.FavoriteProduct{
		UserId:    userId,
		ProductId: uint(id),
	}

	if err := s.repo.createFavorite(newFavorite); err != nil {
		return customerror.GormError(err, "Produk")
	}

	return nil
}

func (s *service) removeFavorite(userId uint, productId string) error {
	id, err := strconv.ParseUint(productId, 10, 32)

	if err != nil {
		return customerror.New("Id produk tidak valid", 400)
	}

	favorite, err := s.repo.checkFavorite(userId, uint(id))

	if err != nil {
		return customerror.GormError(err, "Produk")
	}

	if favorite == nil {
		return customerror.New("Produk tidak ada di dalam list favorit anda", 400)
	}

	if err := s.repo.deleteFavorite(userId, uint(id)); err != nil {
		return customerror.GormError(err, "Produk")
	}

	return nil
}
