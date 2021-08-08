package repository

import (
	"github.com/orensimple/trade-core-app/internal/app/domain"
)

// User is interface of user repository
type User interface {
	Create(u *domain.User) error
	Search(f *domain.User) ([]*domain.User, error)
	Get(f *domain.User) (*domain.User, error)
	Update(f *domain.User) error
	Delete(f *domain.User) error
	Creates(u []domain.User) error
}
