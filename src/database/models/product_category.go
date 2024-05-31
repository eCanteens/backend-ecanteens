package models

type ProductCategory struct {
	PK
	Name string `gorm:"type:varchar(50)" json:"name"`
	Timestamps
}
