package repository

import "github.com/orensimple/trade-core-app/internal/app/domain"

// IOrder is interface of order repository
type IOrder interface {
	Get() domain.Order
	Update(domain.Order)
}
