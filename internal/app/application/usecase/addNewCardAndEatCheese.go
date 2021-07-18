package usecase

import (
	"github.com/orensimple/trade-core-app/internal/app/domain"
	"github.com/orensimple/trade-core-app/internal/app/domain/repository"
)

// AddNewCardAndEatCheese updates payment card
func AddNewCardAndEatCheese(o repository.IOrder) domain.Order {
	order := o.Get()
	o.Update(order)

	return order
}
