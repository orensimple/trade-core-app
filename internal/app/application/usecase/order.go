package usecase

import (
	"github.com/orensimple/trade-core-app/internal/app/adapter/service"
	"github.com/orensimple/trade-core-app/internal/app/domain"
)

// CreateOrder is the UseCase of create new order
func CreateOrder(e service.Order, u *domain.Order) (*domain.Order, error) {
	return e.Create(u)
}

// DeleteOrder is the UseCase of delete order
func DeleteOrder(e service.Order, u *domain.Order) (*domain.Order, error) {
	return e.Delete(u)
}

// FindOrders is the UseCase of find orders
func FindOrders(e service.Order, u *domain.Account) ([]*domain.Order, error) {
	return e.Find(u)
}
