package repository

import "github.com/orensimple/trade-core-app/internal/app/domain"

// IParameter is interface of parameter repository
type IParameter interface {
	Get() domain.Parameter
}
