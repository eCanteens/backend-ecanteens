package product

type feedbackScheme struct {
	IsLiked *bool `json:"is_liked"`
}

type paginationQS struct {
	Page      string `form:"page"`
	Search    string `form:"search" mod:"trim"`
	Limit     string `form:"limit"`
	Order     string `form:"order"`
	Direction string `form:"direction"`
}