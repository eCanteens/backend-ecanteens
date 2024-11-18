package dashboard

import "github.com/eCanteens/backend-ecanteens/src/database/models"

type toggleOpenScheme struct {
	IsOpen bool `binding:"required" json:"is_open"`
}

type summaryDto struct {
	SumTodayTrx    uint `json:"sum_today_trx"`
	SumTodayIncome uint `json:"sum_today_income"`
}

type dashboardDto struct {
	Summary summaryDto   `json:"summary"`
	History models.Order `json:"history"`
}
