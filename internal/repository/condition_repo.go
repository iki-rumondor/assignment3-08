package repository

import "github.com/iki-rumondor/assignment3-08/internal/domain"

type ConditionRepository interface {
	Save(*domain.Condition) (*domain.Condition, error)
}
