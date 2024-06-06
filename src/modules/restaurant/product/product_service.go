package product

import (
	"strconv"

	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/pagination"
	"github.com/eCanteens/backend-ecanteens/src/helpers/upload"
)

func createProductService(user *models.User, body *createProduct) error {
	filepath, err := upload.New(&upload.Option{
		Folder:      "product",
		File:        body.Image,
		NewFilename: strconv.FormatUint(uint64(*user.Id), 10),
	})

	if err != nil {
		return err
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

	if err := create(product); err != nil {
		return err
	}

	return nil
}

func getAllProductService(query *productQs, user *models.User) (*pagination.Pagination, error) {
	var result pagination.Pagination

	if err := findAllProduct(&result, query, user); err != nil {
		return nil, err
	}

	return &result, nil
}
