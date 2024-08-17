package product

import (
	"strconv"

	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/customerror"
	"github.com/eCanteens/backend-ecanteens/src/helpers/pagination"
)

func addFeedbackService(body *feedbackScheme, userId uint, productId string) error {
	id, err := strconv.ParseUint(productId, 10, 32)

	if err != nil {
		return customerror.New("Id produk tidak valid", 400)
	}

	feedbacks, err := checkFeedback(userId, uint(id))

	if err != nil {
		return customerror.GormError(err, "Ulasan")
	}

	if len(*feedbacks) > 0 {
		if err := updateFeedback(*(*feedbacks)[0].Id, body); err != nil {
			return customerror.GormError(err, "Ulasan")
		}

		return nil
	} else {
		feedback := &models.ProductFeedback{
			UserId:    userId,
			ProductId: uint(id),
			IsLike:    *body.IsLike,
		}

		if err := createFeedback(feedback); err != nil {
			return customerror.GormError(err, "Ulasan")
		}

		return nil
	}
}

func removeFeedbackService(userId uint, productId string) error {
	id, err := strconv.ParseUint(productId, 10, 32)

	if err != nil {
		return customerror.New("Id produk tidak valid", 400)
	}

	if err := deleteFeedback(userId, uint(id)); err != nil {
		return customerror.GormError(err, "Ulasan")
	}

	return nil
}

func getFavoriteService(userId uint, query *paginationQS) (*pagination.Pagination[models.Product], error) {
	var result = pagination.New(models.Product{})

	if err := findFavorite(result, userId, query); err != nil {
		return nil, customerror.GormError(err, "Produk")
	}

	return result, nil
}

func addFavoriteService(userId uint, productId string) error {
	id, err := strconv.ParseUint(productId, 10, 32)

	if err != nil {
		return customerror.New("Id produk tidak valid", 400)
	}

	favorites := checkFavorite(userId, uint(id))

	if len(*favorites) > 0 {
		return customerror.New("Produk sudah di dalam list favorit anda", 400)
	}

	favorite := &models.FavoriteProduct{
		UserId:    userId,
		ProductId: uint(id),
	}

	if err := createFavorite(favorite); err != nil {
		return customerror.GormError(err, "Produk")
	}

	return nil
}

func removeFavoriteService(userId uint, productId string) error {
	id, err := strconv.ParseUint(productId, 10, 32)

	if err != nil {
		return customerror.New("Id produk tidak valid", 400)
	}

	favorites := checkFavorite(userId, uint(id))

	if len(*favorites) == 0 {
		return customerror.New("produk tidak ada di dalam list favorit anda", 400)
	}

	if err := deleteFavorite(userId, uint(id)); err != nil {
		return customerror.GormError(err, "Produk")
	}

	return nil
}
