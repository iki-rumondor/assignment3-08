package testing

import (
	"errors"
	"testing"

	"github.com/iki-rumondor/assignment3-08/internal/application"
	"github.com/iki-rumondor/assignment3-08/internal/domain"
	"github.com/iki-rumondor/assignment3-08/internal/testing/mocks"
	"github.com/stretchr/testify/assert"
)

func TestUpdateConditionSuccess(t *testing.T) {
	var mockRepo = mocks.NewConditionRepository(t)
	var condServices = application.NewConditionServices(mockRepo)

	condition := domain.Condition{
		Id:    1,
		Water: 12,
		Wind:  10,
	}

	mockRepo.Mock.On("Save", &condition).Return(
		nil, errors.New("condition with id 1 is not found"),
	)

	cond, err := condServices.UpdateConditon(12, 10)

	assert.Nil(t, cond, "condition should be nil because condition with Id 1 nothing in databse")
	assert.NotNil(t, err, "error should be there because condition with Id 1 nothing in databse")
}

func TestUpdateConditionFailed(t *testing.T) {
	var mockRepo = mocks.NewConditionRepository(t)
	var condServices = application.NewConditionServices(mockRepo)

	condition := domain.Condition{
		Id:    1,
		Water: 12,
		Wind:  10,
	}

	mockRepo.Mock.On("Save", &condition).Return(
		&domain.Condition{
			Id:    1,
			Water: 12,
			Wind:  10,
		}, nil,
	)

	cond, err := condServices.UpdateConditon(12, 10)

	assert.NotNil(t, cond, "condition should be return data condition")
	assert.Nil(t, err, "error should be nil")
}
