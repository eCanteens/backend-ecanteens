package models

type Order struct {
	Id
	UserId    uint   `gorm:"type:bigint" json:"user_id"`
	Quantity  uint   `gorm:"type:int" json:"quantity"`
	Amount    uint   `gorm:"type:int" json:"amount"`
	Notes     string `gorm:"type:varchar(255)" json:"notes"`
	Timestamps

	// Relation
	User    *User    `gorm:"foreignKey:user_id" json:"user,omitempty"`
}