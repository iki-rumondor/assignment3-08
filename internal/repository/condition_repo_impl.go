package repository

import (
	"github.com/iki-rumondor/assignment3-08/internal/domain"
	"gorm.io/gorm"
)

type conditioRepositoryImpl struct {
	db *gorm.DB
}

func NewConditionRepo(db *gorm.DB) ConditionRepository {
	var repo = conditioRepositoryImpl{
		db: db,
	}

	return &repo
}

func (r *conditioRepositoryImpl) Save(condition *domain.Condition) (*domain.Condition, error) {

	if err := r.db.Save(&condition).Error; err != nil {
		return nil, err
	}

	return condition, nil
}
