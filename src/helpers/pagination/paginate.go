package pagination

import (
	"strconv"

	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"gorm.io/gorm"
)

func Paginate(value interface{}, pagination *models.Pagination, params *Params) func(db *gorm.DB) *gorm.DB {
	if params.Page == "" || params.Page == nil {
		pagination.Meta.CurrentPage = 1
	} else {
		switch v := params.Page.(type) {
		case string:
			pagination.Meta.CurrentPage, _ = strconv.Atoi(v)
		case int:
			pagination.Meta.CurrentPage = v
		}
	}

	if params.Limit == "" || params.Limit == nil {
		pagination.Meta.PerPage = 10
	} else {
		switch v := params.Limit.(type) {
		case string:
			pagination.Meta.PerPage, _ = strconv.Atoi(v)
		case int:
			pagination.Meta.PerPage = v
		}
	}

	var totalData int64
	config.DB.Model(value).Count(&totalData)
	pagination.Meta.Total = totalData

	offset := (pagination.Meta.CurrentPage - 1) * pagination.Meta.PerPage

	pagination.Data = value

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(pagination.Meta.PerPage)
	}
}
