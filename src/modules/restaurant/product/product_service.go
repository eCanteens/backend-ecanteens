package product

import (
	"strconv"

	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/customerror"
	"github.com/eCanteens/backend-ecanteens/src/helpers/pagination"
	"github.com/eCanteens/backend-ecanteens/src/helpers/upload"
)

type Service interface {
	createProduct(user *models.User, body *createProduct) error
	getAllProducts(query *productQs, user *models.User) (*getProductsResponse, error)
	updateProduct(user *models.User, body *updateProduct, id string) error
	deleteProduct(user *models.User, productId string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) createProduct(user *models.User, body *createProduct) error {
	filepath, err := upload.New(&upload.Option{
		Folder:      "product",
		File:        body.Image,
		NewFilename: strconv.FormatUint(uint64(*user.Id), 10),
	})

	if err != nil {
		return customerror.New("Gagal saat menyimpan file", 500)
	}

	product := &models.Product{
		RestaurantId: *user.Restaurant.Id,
		Image:        filepath.Url,
		Name:         body.Name,
		Price:        body.Price,
		Stock:        body.Stock,
		Description:  body.Description,
		CategoryId:   body.CategoryId,
	}

	if err := s.repo.create(product); err != nil {
		return customerror.GormError(err, "Menu")
	}

	return nil
}

func (s *service) getAllProducts(query *productQs, user *models.User) (*getProductsResponse, error) {
	var categories []models.ProductCategory

	var responseDto getProductsResponse
	responseDto.Meta.Categories = []*categoryDTO{};
	responseDto.Data = []*categoryProductsDTO{}

	if err := s.repo.findProductCategories(&categories, query.CategoryId); err != nil {
		return nil, err
	}

	for _, category := range categories {
		var categoryDto = categoryDTO{
			Id:   *category.Id,
			Name: category.Name,
		}

		var result = pagination.New(models.Product{})

		if err := s.repo.find(result, &query.paginationQS, *category.Id, *user.Restaurant.Id); err != nil {
			return nil, customerror.GormError(err, "Produk")
		}

		if len(*result.Data) > 0 || query.CategoryId == strconv.Itoa(int(*category.Id)) {
			responseDto.Meta.Categories = append(responseDto.Meta.Categories, &categoryDto)
			responseDto.Data = append(responseDto.Data, &categoryProductsDTO{
				Category:   &categoryDto,
				Pagination: result,
			})
		}
	}

	return &responseDto, nil
}

func (s *service) updateProduct(user *models.User, body *updateProduct, id string) error {
	product := &models.Product{
		RestaurantId: *user.Restaurant.Id,
		Name:         body.Name,
		Price:        body.Price,
		Stock:        body.Stock,
		Description:  body.Description,
		CategoryId:   body.CategoryId,
	}

	if body.Image != nil {
		filepath, err := upload.New(&upload.Option{
			Folder:      "product",
			File:        body.Image,
			NewFilename: strconv.FormatUint(uint64(*user.Id), 10),
		})

		if err != nil {
			return customerror.New("Gagal saat menyimpan file", 500)
		}

		product.Image = filepath.Url
	}

	if err := s.repo.update(product, id); err != nil {
		return customerror.GormError(err, "Menu")
	}

	return nil
}

func (s *service) deleteProduct(user *models.User, productId string) error {
	id, err := strconv.ParseUint(productId, 10, 32)

	if err != nil {
		return customerror.New("Id menu tidak valid", 400)
	}

	if err := s.repo.delete(uint(id), *user.Restaurant.Id); err != nil {
		return customerror.GormError(err, "Menu")
	}

	return nil
}
