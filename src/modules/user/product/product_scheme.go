package product

type FeedbackScheme struct {
	IsLike *bool	`binding:"required" json:"is_like"`
}