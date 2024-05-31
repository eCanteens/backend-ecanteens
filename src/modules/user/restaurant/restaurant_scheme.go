package restaurant

type paginationQS struct {
	Page      string `form:"page"`
	Search    string `form:"search" mod:"trim"`
	Limit     string `form:"limit"`
	Order     string `form:"order"`
	Direction string `form:"direction"`
}

type reviewQS struct {
	Filter string `form:"filter"`
}
