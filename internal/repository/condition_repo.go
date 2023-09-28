package repository

import "github.com/iki-rumondor/assignment3-GLNG-KS-08-08/internal/domain"

type ConditionRepository interface {
	Save(*domain.Condition) (*domain.Condition, error)
}