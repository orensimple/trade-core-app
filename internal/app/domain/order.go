package domain

// Order is an order which has id, payment and user info.
type Order struct {
	ID   string
	User User
}
