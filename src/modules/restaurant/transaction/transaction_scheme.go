package transaction

type getOrderQS struct {
	Page      string `form:"page"`
	Limit     string `form:"limit"`
	Order     string `form:"order"`
	Direction string `form:"direction"`
	Filter    string `form:"filter"`
}