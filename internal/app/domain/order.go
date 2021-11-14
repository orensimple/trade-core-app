package domain

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Order is the model of Order
type Order struct {
	ID           uuid.UUID       `form:"id" json:"id"`
	AccountID    uuid.UUID       `form:"account_id" json:"account_id"`
	InstrumentID uuid.UUID       `form:"instrument_id" json:"instrument_id"`
	Type         string          `form:"type" json:"type"`
	Price        decimal.Decimal `form:"price" json:"price"`
	Volume       int             `form:"volume" json:"volume"`
	Status       string          `form:"status" json:"status"`
}

func (o *Order) Bind(*http.Request) error {
	return nil
}
