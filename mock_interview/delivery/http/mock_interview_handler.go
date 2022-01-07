package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/madjiebimaa/monim/domain"
	"github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"
)

type ResponseErr struct {
	Message string `json:"message"`
}

type MockInterviewHandler struct {
	MockInterviewUsecase domain.MockInterviewUsecase
}

func NewMockInterviewHandler(r *gin.Engine, MockInterviewUsecase domain.MockInterviewUsecase) {
	handler := &MockInterviewHandler{
		MockInterviewUsecase,
	}

	r.GET("/api/mock_interviews/:id", handler.GetByID)
	r.PATCH("/api/mock_interviews", handler.Update)
	r.DELETE("/api/mock_interviews/:id", handler.Delete)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)

	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func isRequestValid(mi *domain.MockInterview) (bool, error) {
	validate := validator.New()
	err := validate.Struct(mi)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (m *MockInterviewHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	ctx := c.Request.Context()
	mi, err := m.MockInterviewUsecase.GetByID(ctx, id)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseErr{
			Message: err.Error(),
		})
	}

	c.JSON(http.StatusOK, mi)
}

func (m *MockInterviewHandler) Update(c *gin.Context) {
	var mi domain.MockInterview
	if err := c.BindJSON(&mi); err != nil {
		c.JSON(http.StatusUnprocessableEntity, ResponseErr{
			Message: err.Error(),
		})
	}

	if ok, err := isRequestValid(&mi); !ok {
		c.JSON(http.StatusBadRequest, ResponseErr{
			Message: err.Error(),
		})
	}

	ctx := c.Request.Context()
	if err := m.MockInterviewUsecase.Update(ctx, &mi); err != nil {
		c.JSON(getStatusCode(err), ResponseErr{
			Message: err.Error(),
		})
	}

	c.JSON(http.StatusNoContent, nil)
}

func (m *MockInterviewHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	ctx := c.Request.Context()
	err := m.MockInterviewUsecase.Delete(ctx, id)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseErr{
			Message: err.Error(),
		})
	}

	c.JSON(http.StatusNoContent, nil)
}
