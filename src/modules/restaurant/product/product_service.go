package product

import (
	"strconv"

	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/upload"
	"github.com/gin-gonic/gin"
)

func createProductService(ctx *gin.Context, user *models.User, body *createProduct) error {
	filepath := upload.New(&upload.Option{
		Folder:      "product",
		Filename:    body.Image.Filename,
		NewFilename: strconv.FormatUint(uint64(*user.Id), 10),
	})

	product := &models.Product{
		RestaurantId: *user.Restaurant.Id,
		Image:       filepath.Url,
		Name:        body.Name,
		Price:       body.Price,
		Stock:       body.Stock,
		Description: body.Description,
		CategoryId:  body.CategoryId,
	}

	if err := create(product); err != nil {
		return err
	}

	return nil
}
