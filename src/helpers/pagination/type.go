package pagination

import "gorm.io/gorm"

type Params struct {
	Query     *gorm.DB
	Page      interface{}
	Limit     interface{}
	Order     string
	Direction string
}
