package repository

import (
	"github.com/orensimple/trade-core-app/internal/app/domain"
)

// Account is interface of account repository
type Account interface {
	Create(u *domain.Account) error
	Get(f *domain.Account) (*domain.Account, error)
	Update(f *domain.Account) error
	Delete(f *domain.Account) error
}
