package usecase

import (
	"github.com/orensimple/trade-core-app/internal/app/application/service"
	"github.com/orensimple/trade-core-app/internal/app/domain/valueobject"
)

// OhlcArgs are arguments of Ohlc usecase
type OhlcArgs struct {
	E service.IExchange
	P valueobject.Pair
	T valueobject.Timeunit
}

// Ohlc is the usecase of getting open, high, low, and close
func Ohlc(a OhlcArgs) []valueobject.CandleStick {
	return a.E.Ohlc(a.P, a.T)
}
