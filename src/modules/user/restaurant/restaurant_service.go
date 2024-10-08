package restaurant

import (
	"strconv"

	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/customerror"
	"github.com/eCanteens/backend-ecanteens/src/helpers/pagination"
)

type Service interface {
	getFavorite(userId uint, query *paginationQS) (*pagination.Pagination[models.Restaurant], error)
	getAll(query *paginationQS) (*pagination.Pagination[models.Restaurant], error)
	getReviews(id string, query *reviewQS) (*[]models.Review, error)
	getDetail(id string) (*models.Restaurant, error)
	getRestosProducts(id string, query *getProductsQS) (*GetProductsResponse, error)
	addFavorite(userId uint, restaurantId string) error
	removeFavorite(userId uint, restaurantId string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) getFavorite(userId uint, query *paginationQS) (*pagination.Pagination[models.Restaurant], error) {
	var result = pagination.New(models.Restaurant{})

	if err := s.repo.findFavorite(result, userId, query); err != nil {
		return nil, customerror.GormError(err, "Restoran")
	}

	return result, nil
}

func (s *service) getAll(query *paginationQS) (*pagination.Pagination[models.Restaurant], error) {
	var result = pagination.New(models.Restaurant{})

	if err := s.repo.find(result, query); err != nil {
		return nil, customerror.GormError(err, "Restoran")
	}

	return result, nil
}

func (s *service) getReviews(id string, query *reviewQS) (*[]models.Review, error) {
	var reviews []models.Review

	if err := s.repo.findReviews(&reviews, id, query); err != nil {
		return nil, customerror.GormError(err, "Restoran")
	}

	return &reviews, nil
}

func (s *service) getDetail(id string) (*models.Restaurant, error) {
	var restaurant models.Restaurant

	if err := s.repo.findOne(&restaurant, id); err != nil {
		return nil, customerror.GormError(err, "Restoran")
	}

	return &restaurant, nil
}

func (s *service) getRestosProducts(id string, query *getProductsQS) (*GetProductsResponse, error) {
	var categories []models.ProductCategory
	var responseDto GetProductsResponse

	if err := s.repo.findProductCategories(&categories, query.CategoryId); err != nil {
		return nil, err
	}

	for _, category := range categories {
		var categoryDto = productCategoryDTO{
			Id:   *category.Id,
			Name: category.Name,
		}

		var result = pagination.New(models.Product{})

		if err := s.repo.findRestosProducts(result, id, &query.paginationQS, *category.Id); err != nil {
			return nil, customerror.GormError(err, "Restoran")
		}

		if len(*result.Data) > 0 {
			responseDto.Meta.Categories = append(responseDto.Meta.Categories, &categoryDto)
			responseDto.Data = append(responseDto.Data, &CategoryProductsDTO{
				Category:   &categoryDto,
				Pagination: result,
			})
		}
	}

	return &responseDto, nil
}

func (s *service) addFavorite(userId uint, restaurantId string) error {
	id, err := strconv.ParseUint(restaurantId, 10, 32)
	if err != nil {
		return customerror.New("Id restoran tidak valid", 400)
	}

	favorites := s.repo.checkFavorite(userId, uint(id))

	if len(*favorites) > 0 {
		return customerror.New("Restoran sudah di dalam list favorit anda", 400)
	}

	favorite := &models.FavoriteRestaurant{
		UserId:       userId,
		RestaurantId: uint(id),
	}

	if err := s.repo.createFavorite(favorite); err != nil {
		return customerror.GormError(err, "Restoran")
	}

	return nil
}

func (s *service) removeFavorite(userId uint, restaurantId string) error {
	id, err := strconv.ParseUint(restaurantId, 10, 32)
	if err != nil {
		return customerror.New("Id restoran tidak valid", 400)
	}

	favorites := s.repo.checkFavorite(userId, uint(id))

	if len(*favorites) == 0 {
		return customerror.New("Restoran tidak ada di dalam list favorit anda", 400)
	}

	if err := s.repo.deleteFavorite(userId, uint(id)); err != nil {
		return customerror.GormError(err, "Restoran")
	}

	return nil
}
