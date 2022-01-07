package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/madjiebimaa/monim/domain"
	"github.com/madjiebimaa/monim/domain/mocks"
	"github.com/madjiebimaa/monim/mock_interview/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetByID(t *testing.T) {
	mockMIRepo := new(mocks.MockInterviewRepository)
	mockMI := domain.MockInterview{
		ID:     uuid.NewString(),
		MeetAt: time.Now(),
		User: domain.User{
			ID: uuid.NewString(),
		},
	}

	t.Run("success get mock interview", func(t *testing.T) {
		mockMIRepo.On("GetByID", mock.Anything, mock.AnythingOfType("string")).Return(mockMI, nil).Once()

		mockUserRepo := new(mocks.UserRepository)
		mockUserRepo.On("GetByID", mock.Anything, mock.AnythingOfType("string")).Return(mockMI.User, nil).Once()

		u := usecase.NewMockInterviewUsecase(mockMIRepo, mockUserRepo, 2*time.Second)

		mi, err := u.GetByID(context.TODO(), mockMI.ID)

		assert.NoError(t, err)
		assert.NotNil(t, mi)
		assert.Equal(t, mockMI.ID, mi.ID)
		assert.Equal(t, mockMI.User.ID, mi.User.ID)
		mockMIRepo.AssertExpectations(t)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("fail because mock interview db error", func(t *testing.T) {
		mockMIRepo.On("GetByID", mock.Anything, mock.AnythingOfType("string")).Return(domain.MockInterview{}, domain.ErrUnexpected).Once()

		mockUserRepo := new(mocks.UserRepository)

		u := usecase.NewMockInterviewUsecase(mockMIRepo, mockUserRepo, 2*time.Second)

		mi, err := u.GetByID(context.TODO(), mockMI.ID)

		assert.Error(t, err)
		assert.Equal(t, domain.MockInterview{}, mi)
		mockMIRepo.AssertExpectations(t)
	})

	t.Run("fail because user db error", func(t *testing.T) {
		mockMIRepo.On("GetByID", mock.Anything, mock.AnythingOfType("string")).Return(mockMI, nil).Once()

		mockUserRepo := new(mocks.UserRepository)
		mockUserRepo.On("GetByID", mock.Anything, mock.AnythingOfType("string")).Return(domain.User{}, domain.ErrUnexpected).Once()

		u := usecase.NewMockInterviewUsecase(mockMIRepo, mockUserRepo, 2*time.Second)

		mi, err := u.GetByID(context.TODO(), mockMI.ID)

		assert.Error(t, err)
		assert.Equal(t, domain.User{}, mi.User)
		mockMIRepo.AssertExpectations(t)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	mockMIRepo := new(mocks.MockInterviewRepository)
	mockMI := domain.MockInterview{
		ID:     uuid.NewString(),
		MeetAt: time.Now(),
	}

	t.Run("success update mock interview", func(t *testing.T) {
		mockMIRepo.On("Update", mock.Anything, &mockMI).Return(nil).Once()

		mockUserRepo := new(mocks.UserRepository)
		u := usecase.NewMockInterviewUsecase(mockMIRepo, mockUserRepo, 2*time.Second)

		err := u.Update(context.TODO(), &mockMI)

		assert.NoError(t, err)
		mockMIRepo.AssertExpectations(t)
	})

}

func TestDelete(t *testing.T) {
	mockMIRepo := new(mocks.MockInterviewRepository)
	mockMI := domain.MockInterview{
		ID:     uuid.NewString(),
		MeetAt: time.Now(),
	}

	t.Run("success delete mock interview", func(t *testing.T) {
		mockMIRepo.On("GetByID", mock.Anything, mock.AnythingOfType("string")).Return(mockMI, nil).Once()
		mockMIRepo.On("Delete", mock.Anything, mock.AnythingOfType("string")).Return(nil).Once()

		mockUserRepo := new(mocks.UserRepository)
		u := usecase.NewMockInterviewUsecase(mockMIRepo, mockUserRepo, 2*time.Second)

		err := u.Delete(context.TODO(), mockMI.ID)

		assert.NoError(t, err)
		mockMIRepo.AssertExpectations(t)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("fail because error in db", func(t *testing.T) {
		mockMIRepo.On("GetByID", mock.Anything, mock.AnythingOfType("string")).Return(domain.MockInterview{}, domain.ErrUnexpected).Once()

		mockUserRepo := new(mocks.UserRepository)
		u := usecase.NewMockInterviewUsecase(mockMIRepo, mockUserRepo, 2*time.Second)

		err := u.Delete(context.TODO(), mockMI.ID)
		assert.Error(t, err)
		mockMIRepo.AssertExpectations(t)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("fail because mock interview not exist", func(t *testing.T) {
		mockMIRepo.On("GetByID", mock.Anything, mock.AnythingOfType("string")).Return(domain.MockInterview{}, nil)

		mockUserRepo := new(mocks.UserRepository)
		u := usecase.NewMockInterviewUsecase(mockMIRepo, mockUserRepo, 2*time.Second)

		err := u.Delete(context.TODO(), mockMI.ID)
		assert.Error(t, err)
		mockMIRepo.AssertExpectations(t)
		mockUserRepo.AssertExpectations(t)
	})
}
