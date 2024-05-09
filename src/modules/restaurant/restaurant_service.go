package restaurant

import (
	"errors"
	"strconv"

	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/pagination"
)

func getFavoriteService(userId uint, query map[string]string) (*[]models.Restaurant, error) {
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

	return &user.FavoriteRestaurant, nil
}

func getAllService(query map[string]string) (*pagination.Pagination, error) {
	var restoPagination pagination.Pagination
	var restaurants []models.Restaurant

	if err := find(&restoPagination, &restaurants, query); err != nil {
		return nil, err
	}

	return &restoPagination, nil
}

func getDetailService(id string) (*models.Restaurant, error) {
	var restaurant models.Restaurant

	if err := findOne(&restaurant, id); err != nil {
		return nil, err
	}

	return &restaurant, nil
}

func getRestosProductsService(id string, query map[string]string) (*pagination.Pagination, error) {
	var productsPagination pagination.Pagination
	var products []models.Product

	if err := findRestosProducts(&productsPagination, &products, id, query); err != nil {
		return nil, err
	}

	return &productsPagination, nil
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

	favorite := &models.Favorite{
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