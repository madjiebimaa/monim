package mocks

import (
	"context"

	"github.com/madjiebimaa/monim/domain"
	"github.com/stretchr/testify/mock"
)

type MockInterviewRepository struct {
	mock.Mock
}

func (_m *MockInterviewRepository) GetByID(ctx context.Context, id string) (domain.MockInterview, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.MockInterview
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.MockInterview); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.MockInterview)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockInterviewRepository) Update(ctx context.Context, mi *domain.MockInterview) error {
	ret := _m.Called(ctx, mi)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.MockInterview) error); ok {
		r0 = rf(ctx, mi)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *MockInterviewRepository) Delete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
