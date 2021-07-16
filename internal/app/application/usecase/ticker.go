package usecase

import (
	"github.com/orensimple/trade-core-app/internal/app/application/service"
	"github.com/orensimple/trade-core-app/internal/app/domain/valueobject"
)

// Ticker is the usecase of getting ticker
func Ticker(e service.IExchange, p valueobject.Pair) valueobject.Ticker {
	return e.Ticker(p)
}
