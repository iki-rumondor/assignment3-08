package application

import (
	"github.com/iki-rumondor/assignment3-GLNG-KS-08-08/internal/adapter/http/response"
	"github.com/iki-rumondor/assignment3-GLNG-KS-08-08/internal/domain"
	"github.com/iki-rumondor/assignment3-GLNG-KS-08-08/internal/repository"
)

type ConditionServices struct {
	Repo repository.ConditionRepository
}

func NewConditionServices(repo repository.ConditionRepository) *ConditionServices {
	services := ConditionServices{
		Repo: repo,
	}

	return &services
}

func (s *ConditionServices) UpdateConditon(water int, wind int) (*response.Condition, error) {

	condtion := domain.Condition{
		Id:    1,
		Water: water,
		Wind:  wind,
	}

	condFromDB, err := s.Repo.Save(&condtion)
	if err != nil {
		return nil, err
	}

	res := response.Condition{
		Water: condFromDB.Water,
		Wind:  condFromDB.Wind,
	}

	return &res, nil
}
