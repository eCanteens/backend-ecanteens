package product

import (
	"errors"
	"strconv"

	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/pagination"
)

func addFeedbackService(body *feedbackScheme, userId uint, productId string) error {
	id, err := strconv.ParseUint(productId, 10, 32)

	if err != nil {
		return err
	}

	feedbacks, err := checkFeedback(userId, uint(id))

	if err != nil {
		return err
	}

	if len(*feedbacks) > 0 {
		return updateFeedback(*(*feedbacks)[0].Id, body)
	} else {
		feedback := &models.ProductFeedback{
			UserId:    userId,
			ProductId: uint(id),
			IsLike:    *body.IsLike,
		}

		return createFeedback(feedback)
	}
}

func removeFeedbackService(userId uint, productId string) error {
	id, err := strconv.ParseUint(productId, 10, 32)

	if err != nil {
		return err
	}

	return deleteFeedback(userId, uint(id))
}

func getFavoriteService(userId uint, query *paginationQS) (*pagination.Pagination[models.Product], error) {
	var result = pagination.New(models.Product{})

	if err := findFavorite(result, userId, query); err != nil {
		return nil, err
	}

	return result, nil
}

func addFavoriteService(userId uint, productId string) error {
	id, err := strconv.ParseUint(productId, 10, 32)
	if err != nil {
		return err
	}

	favorites := checkFavorite(userId, uint(id))

	if len(*favorites) > 0 {
		return errors.New("produk sudah di dalam list favorit anda")
	}

	favorite := &models.FavoriteProduct{
		UserId:    userId,
		ProductId: uint(id),
	}

	if err := createFavorite(favorite); err != nil {
		return err
	}

	return nil
}

func removeFavoriteService(userId uint, productId string) error {
	id, err := strconv.ParseUint(productId, 10, 32)
	if err != nil {
		return err
	}

	favorites := checkFavorite(userId, uint(id))

	if len(*favorites) == 0 {
		return errors.New("produk tidak ada di dalam list favorit anda")
	}

	if err := deleteFavorite(userId, uint(id)); err != nil {
		return err
	}

	return nil
}
