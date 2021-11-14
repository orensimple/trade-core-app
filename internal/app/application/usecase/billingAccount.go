package usecase

import (
	"github.com/orensimple/trade-core-app/internal/app/adapter/service"
	"github.com/orensimple/trade-core-app/internal/app/domain"
)

// CreateBillingAccount is the UseCase of create account in billing
func CreateBillingAccount(e service.Billing, u *domain.User) (*service.BillingCreateResponse, error) {
	return e.Create(u)
}
