package usecase

import (
	"context"
	"time"

	"github.com/madjiebimaa/monim/domain"
)

type mockInterviewUsecase struct {
	mockInterviewRepo domain.MockInterviewRepository
	userRepo          domain.UserRepository
	contextTimeout    time.Duration
}

func NewMockInterviewUsecase(
	mockInterviewRepo domain.MockInterviewRepository,
	userRepo domain.UserRepository,
	contextTimeout time.Duration,
) domain.MockInterviewUsecase {
	return &mockInterviewUsecase{
		mockInterviewRepo,
		userRepo,
		contextTimeout,
	}
}

func (m *mockInterviewUsecase) GetByID(c context.Context, id string) (domain.MockInterview, error) {
	ctx, cancel := context.WithTimeout(c, m.contextTimeout)
	defer cancel()

	mi, err := m.mockInterviewRepo.GetByID(ctx, id)
	if err != nil {
		return domain.MockInterview{}, err
	}

	user, err := m.userRepo.GetByID(ctx, mi.User.ID)
	if err != nil {
		return domain.MockInterview{}, err
	}

	mi.User = user
	return mi, nil
}

func (m *mockInterviewUsecase) Update(c context.Context, mi *domain.MockInterview) error {
	ctx, cancel := context.WithTimeout(c, m.contextTimeout)
	defer cancel()

	mi.UpdatedAt = time.Now()
	return m.mockInterviewRepo.Update(ctx, mi)
}

func (m *mockInterviewUsecase) Delete(c context.Context, id string) error {
	ctx, cancel := context.WithTimeout(c, m.contextTimeout)
	defer cancel()

	existedMI, err := m.mockInterviewRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if existedMI == (domain.MockInterview{}) {
		return domain.ErrNotFound
	}

	return m.mockInterviewRepo.Delete(ctx, id)
}
