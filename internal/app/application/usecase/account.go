package usecase

import (
	"github.com/orensimple/trade-core-app/internal/app/domain"
	"github.com/orensimple/trade-core-app/internal/app/domain/repository"
)

// CreateAccount create new account
func CreateAccount(r repository.Account, u *domain.Account) (*domain.Account, error) {
	err := r.Create(u)

	return u, err
}

// GetAccount find account by filter
func GetAccount(r repository.Account, f *domain.Account) (*domain.Account, error) {
	res, err := r.Get(f)

	return res, err
}

// UpdateAccount update account
func UpdateAccount(r repository.Account, f *domain.Account) error {
	return r.Update(f)
}

// DeleteAccount delete account by id
func DeleteAccount(r repository.Account, f *domain.Account) error {
	return r.Delete(f)
}
