package domain

import "github.com/orensimple/trade-core-app/internal/app/domain/valueobject"

// Order is an order which has id, payment and user info.
type Order struct {
	ID      string
	Payment valueobject.Payment
	User    User
}
