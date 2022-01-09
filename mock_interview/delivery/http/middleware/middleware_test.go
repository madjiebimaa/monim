package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/madjiebimaa/monim/mock_interview/delivery/http/middleware"
)

func TestCORS(t *testing.T) {
	r := gin.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(res)
	m := middleware.InitMiddleware()

	h := m.CORS(gin.HandlerFunc(func(c *gin.Context) {
		c.JSON(http.StatusNoContent, nil)
	}))

	r.ServeHTTP(res, req)

	h(c)

	assert.Equal(t, "*", res.Header().Get("Access-Control-Allow-Origin"))
}
