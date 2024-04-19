package pagination

import (
	"math"
	"strconv"

	"github.com/eCanteens/backend-ecanteens/src/config"
)

func (pagination *Pagination) Paginate(value interface{}, params *Params) error {
	if params.Query == nil {
		params.Query = config.DB
	}

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

	if params.Order == "" {
		params.Order = "created_at"
	}

	if params.Direction == "" {
		params.Direction = "desc"
	}

	var totalData int64
	params.Query.Model(value).Count(&totalData)
	offset := (pagination.Meta.CurrentPage - 1) * pagination.Meta.PerPage

	pagination.Meta.Total = totalData
	pagination.Meta.LastPage = int(math.Ceil(float64(totalData) / float64(pagination.Meta.PerPage)))
	pagination.Data = value

	return params.Query.Offset(offset).Limit(pagination.Meta.PerPage).Order(params.Order + " " + params.Direction).Find(value).Error
}
