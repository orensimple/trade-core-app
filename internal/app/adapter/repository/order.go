package repository

import (
	"errors"

	"github.com/prometheus/common/log"

	"github.com/orensimple/trade-core-app/internal/app/adapter/mysql/model"
	"github.com/orensimple/trade-core-app/internal/app/domain"
	"github.com/orensimple/trade-core-app/internal/app/domain/factory"
	"gorm.io/gorm"
)

// Order is the repository of domain.Order.
type Order struct {
	repo *gorm.DB
}

func NewOrderRepo(db *gorm.DB) Order {
	return Order{repo: db}
}

// Get gets order.
func (o Order) Get() domain.Order {
	var order model.Order
	// Order has Person/Payment relation and Payment has Card relation which has CardBrand relation.
	result := o.repo.Preload("User").Preload("Payment.Card.CardBrand").Find(&order)
	if result.Error != nil {
		log.Error(result.Error)
	}

	orderFactory := factory.Order{}

	return orderFactory.Generate(
		order.User.ID,
		order.User.Email,
		order.User.FirstName,
		order.User.LastName,
		order.User.Male,
		order.User.About,
		order.User.Address,
		order.Payment.Card.ID,
		order.Payment.Card.CardBrand.Brand,
		order.ID,
	)
}

// Update updates order.
func (o Order) Update(order domain.Order) {
	user := model.User{
		ID:        order.User.ID,
		FirstName: order.User.FirstName,
	}

	err := o.repo.Transaction(func(tx *gorm.DB) error {
		var err error
		err = tx.Exec("update users set first_name = ? where id = ?", user.FirstName, user.ID).Error
		if err != nil {
			return errors.New("rollback")
		}

		return nil
	})
	if err != nil {
		log.Error(err)
	}
}
