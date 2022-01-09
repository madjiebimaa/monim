package middleware

import "github.com/gin-gonic/gin"

type GoMiddleware struct{}

func InitMiddleware() *GoMiddleware {
	return &GoMiddleware{}
}

func (m *GoMiddleware) CORS(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		next(c)
	}
}
