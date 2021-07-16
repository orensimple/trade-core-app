package usecase

import (
	"github.com/google/uuid"
	"github.com/orensimple/trade-core-app/internal/app/domain"
	"github.com/orensimple/trade-core-app/internal/app/domain/repository"
	"github.com/orensimple/trade-core-app/internal/app/domain/valueobject"
)

// AddNewCardAndEatCheese updates payment card
func AddNewCardAndEatCheese(o repository.IOrder) domain.Order {
	order := o.Get()
	newCardBrand := valueobject.VISA

	if order.Payment.Card.Brand == valueobject.VISA {
		newCardBrand = valueobject.AMEX
	}

	newCard := valueobject.Card{
		ID:    uuid.New().String(),
		Brand: newCardBrand,
	}
	order.Payment.Card = newCard
	o.Update(order)

	return order
}
