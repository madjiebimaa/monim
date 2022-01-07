package domain

import (
	"context"
	"time"
)

type MockInterview struct {
	ID                  string    `json:"id"`
	Language            string    `json:"language"`
	ProgrammingLanguage string    `json:"programming_language"`
	User                User      `json:"user"`
	MeetAt              time.Time `json:"meet_at"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

type MockInterviewUsecase interface {
	GetByID(ctx context.Context, id string) (MockInterview, error)
	Update(ctx context.Context, mi *MockInterview) error
	Delete(ctx context.Context, id string) error
}

type MockInterviewRepository interface {
	GetByID(ctx context.Context, id string) (MockInterview, error)
	Update(ctx context.Context, mi *MockInterview) error
	Delete(ctx context.Context, id string) error
}
