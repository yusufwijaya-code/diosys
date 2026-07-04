package middlewares

import (
	"fmt"

	"portfolio-api/base/helpers/error_helper"
	"portfolio-api/base/helpers/http_helper"

	"github.com/gin-gonic/gin"
)

// PanicMiddleware recovers from panics and returns a structured 500 response.
func (m Middleware) PanicMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if recovered := recover(); recovered != nil {
				http_helper.HttpErrorResponse(c, error_helper.Internal(fmt.Errorf("%v", recovered)))
			}
		}()
		c.Next()
	}
}
