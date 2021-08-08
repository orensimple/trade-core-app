package repository

import (
	"errors"
	"fmt"

	"github.com/orensimple/trade-core-app/internal/app/domain"
	"gorm.io/gorm"
)

// User is the repository of domain.User
type User struct {
	repo *gorm.DB
}

func NewUserRepo(db *gorm.DB) User {
	return User{repo: db}
}

// Create create new user
func (u User) Create(user *domain.User) error {
	err := u.repo.Create(user)
	if err != nil {
		return errors.New("can't create new user")
	}

	return nil
}

// Get get user by filter
func (u User) Get(f *domain.User) (*domain.User, error) {
	out := new(domain.User)

	err := u.repo.Where(f).Take(out).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}

		return nil, errors.New("failed get user")
	}

	return out, nil
}

// Update user info by id
func (u User) Update(user *domain.User) error {
	return u.repo.Debug().Save(&user).Error
}

// Delete user by id
func (u User) Delete(f *domain.User) error {
	return u.repo.Delete(&f).Error
}

// Search get user by firstname and lastname
func (u User) Search(f *domain.User) ([]*domain.User, error) {
	out := make([]*domain.User, 0)

	sql := "select users.* from users where last_name like ? and first_name like ? order by id limit 100"
	err := u.repo.Raw(sql, fmt.Sprintf("%s%%", f.LastName), fmt.Sprintf("%s%%", f.FirstName)).Find(&out).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return out, nil
		}

		return nil, errors.New("failed search users")
	}

	return out, nil
}

// Creates create new users
func (u User) Creates(users []domain.User) error {
	err := u.repo.Create(&users)
	if err != nil {
		return errors.New("can't creates new user")
	}

	return nil
}
