package restaurant

import (
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/pagination"
)

func GetFavoriteRestoService(userId uint, query map[string]string) (*[]models.Restaurant, error) {
	var user models.User

	if query["order"] == "" {
		query["order"] = "created_at"
	}

	if query["direction"] == "" {
		query["direction"] = "desc"
	}

	if err := FindFavoriteResto(&user, userId, query); err != nil {
		return nil, err
	}

	return &user.FavoriteRestaurant, nil
}

func GetAllRestoService(query map[string]string) (*pagination.Pagination, error) {
	var restoPagination pagination.Pagination
	var restaurants []models.Restaurant

	if err := FindResto(&restoPagination, &restaurants, query); err != nil {
		return nil, err
	}

	return &restoPagination, nil
}

func GetDetailRestoService(id string) (*models.Restaurant, error) {
	var restaurant models.Restaurant

	if err := FindOneResto(&restaurant, id); err != nil {
		return nil, err
	}

	return &restaurant, nil
}

func GetRestosProductsService(id string, query map[string]string) (*pagination.Pagination, error) {
	var productsPagination pagination.Pagination
	var products []models.Product

	if err := FindRestosProducts(&productsPagination, &products, id, query); err != nil {
		return nil, err
	}

	return &productsPagination, nil
}