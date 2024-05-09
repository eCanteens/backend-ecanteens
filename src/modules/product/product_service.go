package product

import (
	"errors"
	"strconv"

	"github.com/eCanteens/backend-ecanteens/src/database/models"
)

func addFeedbackService(body *FeedbackScheme, userId uint, productId string) error {
	id, err := strconv.ParseUint(productId, 10, 32)

	if err != nil {
		return err
	}

	feedbacks, err := checkFeedback(userId, uint(id))

	if err != nil {
		return err
	}

	if len(*feedbacks) > 0 {
		return updateFeedback(*(*feedbacks)[0].Id.Id, body)
	} else {
		feedback := &models.ProductFeedback{
			UserId: userId,
			ProductId: uint(id),
			IsLike: *body.IsLike,
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

func getFavoriteService(userId uint, query map[string]string) (*[]models.Product, error) {
	var user models.User

	if query["order"] == "" {
		query["order"] = "created_at"
	}

	if query["direction"] == "" {
		query["direction"] = "desc"
	}

	if err := findFavorite(&user, userId, query); err != nil {
		return nil, err
	}

	return &user.FavoriteProducts, nil
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
		UserId: userId,
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