package restaurant

import (
	"errors"
	"strconv"

	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/pagination"
)

func getFavoriteService(userId uint, query *paginationQS) (*pagination.Pagination[models.Restaurant], error) {
	var result = pagination.New(models.Restaurant{})

	if err := findFavorite(result, userId, query); err != nil {
		return nil, err
	}

	return result, nil
}

func getAllService(query *paginationQS) (*pagination.Pagination[models.Restaurant], error) {
	var result = pagination.New(models.Restaurant{})

	if err := find(result, query); err != nil {
		return nil, err
	}

	return result, nil
}

func getReviewsService(id string, query *reviewQS) (*[]models.Review, error) {
	var reviews []models.Review

	if err := findReviews(&reviews, id, query); err != nil {
		return nil, err
	}

	return &reviews, nil
}

func getDetailService(id string) (*models.Restaurant, error) {
	var restaurant models.Restaurant

	if err := findOne(&restaurant, id); err != nil {
		return nil, err
	}

	return &restaurant, nil
}

func getRestosProductsService(id string, query *paginationQS) (*pagination.Pagination[models.Product], error) {
	var result = pagination.New(models.Product{})

	if err := findRestosProducts(result, id, query); err != nil {
		return nil, err
	}

	return result, nil
}

func addFavoriteService(userId uint, restaurantId string) error {
	id, err := strconv.ParseUint(restaurantId, 10, 32)
	if err != nil {
		return err
	}

	favorites := checkFavorite(userId, uint(id))

	if len(*favorites) > 0 {
		return errors.New("restoran sudah di dalam list favorit anda")
	}

	favorite := &models.FavoriteRestaurant{
		UserId: userId,
		RestaurantId: uint(id),
	}

	if err := createFavorite(favorite); err != nil {
		return err
	}

	return nil
}

func removeFavoriteService(userId uint, restaurantId string) error {
	id, err := strconv.ParseUint(restaurantId, 10, 32)
	if err != nil {
		return err
	}

	favorites := checkFavorite(userId, uint(id))

	if len(*favorites) == 0 {
		return errors.New("restoran tidak ada di dalam list favorit anda")
	}

	if err := deleteFavorite(userId, uint(id)); err != nil {
		return err
	}

	return nil
}