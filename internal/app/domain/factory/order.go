package factory

import (
	"github.com/google/uuid"
	"github.com/orensimple/trade-core-app/internal/app/domain"
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
	return domain.Order{
		ID:   orderID,
		User: user,
	}
}
