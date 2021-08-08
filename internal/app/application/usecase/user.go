package usecase

import (
	"github.com/brianvoe/gofakeit"
	"github.com/google/uuid"
	"github.com/orensimple/trade-core-app/internal/app/domain"
	"github.com/orensimple/trade-core-app/internal/app/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

const BatchInsertSize = 1000

// CreateUser create new user
func CreateUser(r repository.User, u *domain.User) (*domain.User, error) {
	err := r.Create(u)

	return u, err
}

// SearchUsers by firstname and lastname
func SearchUsers(r repository.User, f *domain.User) ([]*domain.User, error) {
	u, err := r.Search(f)

	return u, err
}

// GetUser find user by filter
func GetUser(r repository.User, f *domain.User) (*domain.User, error) {
	res, err := r.Get(f)

	return res, err
}

// UpdateUser update user
func UpdateUser(r repository.User, f *domain.User) error {
	return r.Update(f)
}

// DeleteUser delete user by id
func DeleteUser(r repository.User, f *domain.User) error {
	return r.Delete(f)
}

// CreateUsersMock create new mock users for tests
func CreateUsersMock(r repository.User, count int) error {
	var err error

	for i := 0; i < count/BatchInsertSize; i++ {
		go func() {
			users := make([]domain.User, 0, BatchInsertSize)
			for j := 0; j < BatchInsertSize; j++ {
				users = append(users, generateUser())
			}

			err = r.Creates(users)
		}()
	}

	return err
}

func generateUser() domain.User {
	hash, _ := bcrypt.GenerateFromPassword([]byte(gofakeit.Word()), bcrypt.MinCost)

	newUser := domain.User{
		ID:        uuid.New(),
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
		Password:  string(hash),
		Address:   gofakeit.City(),
		About:     gofakeit.Company(),
		Male:      gofakeit.Bool(),
	}

	return newUser
}
