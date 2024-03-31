package models

type ProductCategory struct {
	Id
	Name string `gorm:"type:varchar(50)" json:"name"`
	Timestamps
}
