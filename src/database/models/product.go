package models

import (
	"github.com/eCanteens/backend-ecanteens/src/config"

	"gorm.io/gorm"
)

type FeedbackResult struct {
	Like    int
	Dislike int
}

type Product struct {
	Id
	RestaurantId uint   `gorm:"type:bigint" json:"restaurant_id"`
	Name         string `gorm:"type:varchar(100)" json:"name"`
	Description  string `gorm:"type:text" json:"description"`
	Image        string `gorm:"type:varchar(255)" json:"image"`
	CategoryId   uint   `gorm:"type:bigint" json:"category_id"`
	Price        uint   `gorm:"type:int" json:"price"`
	Stock        uint   `gorm:"type:int" json:"stock"`
	Sold         uint   `gorm:"type:int" json:"sold"`
	Timestamps

	// Relations
	Restaurant *Restaurant      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:restaurant_id" json:"restaurant,omitempty"`
	Category   *ProductCategory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:category_id" json:"category,omitempty"`

	// Extra
	Feedback *FeedbackResult `gorm:"-" json:"feedback"`
}

func (p *Product) AfterFind(tx *gorm.DB) (err error) {
	var result FeedbackResult

	rows, _ := config.DB.Table("product_feedbacks").
		Select("CASE WHEN is_like = true THEN 'like' ELSE 'dislike' END AS feedback_type, COUNT(*) AS count").
		Where("product_id = ?", p.Id.Id).
		Group("feedback_type").
		Rows()

	defer rows.Close()

	for rows.Next() {
        var feedbackType string
        var count int
        if err := rows.Scan(&feedbackType, &count); err != nil {
            return err
        }
        if feedbackType == "like" {
            result.Like = count
        } else {
            result.Dislike = count
        }
    }

	p.Feedback = &result

	return
}
