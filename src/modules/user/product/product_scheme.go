package product

type feedbackScheme struct {
	IsLike *bool `binding:"required" json:"is_like"`
}

type paginationQS struct {
	Page      string `form:"page"`
	Search    string `form:"search" mod:"trim"`
	Limit     string `form:"limit"`
	Order     string `form:"order"`
	Direction string `form:"direction"`
}