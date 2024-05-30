package product

type feedbackScheme struct {
	IsLike *bool `binding:"required" json:"is_like"`
}
