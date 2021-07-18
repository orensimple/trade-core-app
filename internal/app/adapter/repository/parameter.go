package repository

import (
	"github.com/orensimple/trade-core-app/internal/app/adapter/mysql/model"
	"github.com/orensimple/trade-core-app/internal/app/domain"
	"github.com/prometheus/common/log"
	"gorm.io/gorm"
)

// Parameter is the repository of domain.Parameter.
type Parameter struct {
	repo *gorm.DB
}

func NewParameterRepo(db *gorm.DB) Parameter {
	return Parameter{repo: db}
}

// Get gets parameter.
func (r Parameter) Get() domain.Parameter {
	var param model.Parameter
	result := r.repo.First(&param, 1)
	if result.Error != nil {
		log.Error(result.Error)
	}

	return domain.Parameter{
		Funds: param.Funds,
		Btc:   param.Btc,
	}
}
