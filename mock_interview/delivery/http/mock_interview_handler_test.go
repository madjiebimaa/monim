package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/madjiebimaa/monim/domain"
	"github.com/madjiebimaa/monim/domain/mocks"
	miHttp "github.com/madjiebimaa/monim/mock_interview/delivery/http"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	var mockMI domain.MockInterview
	err := faker.FakeData(&mockMI)
	assert.NoError(t, err)

	mockUCase := new(mocks.MockInterviewUsecase)

	t.Run("success get MI", func(t *testing.T) {
		mockUCase.On("GetByID", mock.Anything, mockMI.ID).Return(mockMI, nil).Once()

		rec := httptest.NewRecorder()
		r := gin.New()
		miHttp.NewMockInterviewHandler(r, mockUCase)

		req, err := http.NewRequest(http.MethodGet, "/api/mock_interviews/"+mockMI.ID, nil)
		assert.NoError(t, err)

		r.ServeHTTP(rec, req)

		respBody, err := json.Marshal(mockMI)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, respBody, rec.Body.Bytes())
		mockUCase.AssertExpectations(t)
	})

	t.Run("fail because error in MI use case", func(t *testing.T) {
		mockUCase.On("GetByID", mock.Anything, mockMI.ID).Return(domain.MockInterview{}, domain.ErrNotFound).Once()

		rec := httptest.NewRecorder()
		r := gin.New()
		miHttp.NewMockInterviewHandler(r, mockUCase)

		req, err := http.NewRequest(http.MethodGet, "/api/mock_interviews/"+mockMI.ID, nil)
		assert.NoError(t, err)

		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)
		mockUCase.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	gin.SetMode(gin.TestMode)

	var mockMI domain.MockInterview
	err := faker.FakeData(&mockMI)
	assert.NoError(t, err)

	mockUCase := new(mocks.MockInterviewUsecase)

	t.Run("success update mock interview", func(t *testing.T) {
		mockUCase.On("Update", mock.Anything, mock.AnythingOfType("*domain.MockInterview")).Return(nil).Once()

		rec := httptest.NewRecorder()
		r := gin.New()
		miHttp.NewMockInterviewHandler(r, mockUCase)

		j, err := json.Marshal(mockMI)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPatch, "/api/mock_interviews", strings.NewReader(string(j)))
		assert.NoError(t, err)

		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNoContent, rec.Code)
		mockUCase.AssertExpectations(t)
	})

	t.Run("fail because there's error in MI use case", func(t *testing.T) {
		mockUCase.On("Update", mock.Anything, mock.AnythingOfType("*domain.MockInterview")).Return(domain.ErrInternalServerError).Once()

		rec := httptest.NewRecorder()
		r := gin.New()
		miHttp.NewMockInterviewHandler(r, mockUCase)

		j, err := json.Marshal(mockMI)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPatch, "/api/mock_interviews", strings.NewReader(string(j)))
		assert.NoError(t, err)

		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockUCase.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	gin.SetMode(gin.TestMode)

	id := uuid.NewString()

	mockUCase := new(mocks.MockInterviewUsecase)

	t.Run("success delete MI", func(t *testing.T) {
		mockUCase.On("Delete", mock.Anything, id).Return(nil).Once()

		rec := httptest.NewRecorder()
		r := gin.New()
		miHttp.NewMockInterviewHandler(r, mockUCase)

		req, err := http.NewRequest(http.MethodDelete, "/api/mock_interviews/"+id, nil)
		assert.NoError(t, err)

		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNoContent, rec.Code)
		mockUCase.AssertExpectations(t)
	})

	t.Run("fail because error in MI use case", func(t *testing.T) {
		mockUCase.On("Delete", mock.Anything, id).Return(domain.ErrInternalServerError).Once()

		rec := httptest.NewRecorder()
		r := gin.New()
		miHttp.NewMockInterviewHandler(r, mockUCase)

		req, err := http.NewRequest(http.MethodDelete, "/api/mock_interviews/"+id, nil)
		assert.NoError(t, err)

		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockUCase.AssertExpectations(t)
	})
}
