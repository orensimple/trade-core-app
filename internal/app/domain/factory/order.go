package factory

import (
	"github.com/google/uuid"
	"github.com/orensimple/trade-core-app/internal/app/domain"
	"github.com/orensimple/trade-core-app/internal/app/domain/valueobject"
)

// Order is the factory of domain.Order
type Order struct{}

// Generate generates domain.Order from primitives
// TODO to dto
func (of Order) Generate(
	userID uuid.UUID,
	email string,
	firstName string,
	lastName string,
	male bool,
	about string,
	address string,
	cardID string,
	brand string,
	orderID string,
) domain.Order {
	user := domain.User{
		ID:        userID,
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Male:      male,
		About:     about,
		Address:   address,
	}
	cardBrand := valueobject.ConvertToCardBrand(brand)
	card := valueobject.Card{
		ID:    cardID,
		Brand: cardBrand,
	}
	payment := valueobject.Payment{
		Card: card,
	}
	return domain.Order{
		ID:      orderID,
		Payment: payment,
		User:    user,
	}
}
