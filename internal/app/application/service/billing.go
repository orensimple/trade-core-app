package service

import (
	"github.com/orensimple/trade-core-app/internal/app/adapter/service"
	"github.com/orensimple/trade-core-app/internal/app/domain"
)

// Billing is interface of billing app http
type Billing interface {
	Create(u domain.User) service.BillingCreateResponse
}
