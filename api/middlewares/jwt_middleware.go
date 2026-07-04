package middlewares

import (
	"strings"

	"portfolio-api/base/helpers/context_helper"
	"portfolio-api/base/helpers/error_helper"
	"portfolio-api/base/helpers/http_helper"
	"portfolio-api/base/helpers/jwt_helper"

	"portfolio-api/constants"

	"github.com/gin-gonic/gin"
)

// JwtMiddleware validates the bearer token and stores the identity in the context.
func (m Middleware) JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			http_helper.HttpErrorResponse(c, error_helper.Unauthorized("missing or invalid authorization header"))
			return
		}

		tokenString := strings.TrimPrefix(header, "Bearer ")
		claims, err := jwt_helper.ParseToken(tokenString, m.Config.JwtSecret)
		if err != nil {
			http_helper.HttpErrorResponse(c, error_helper.Unauthorized("invalid or expired token"))
			return
		}

		c.Set(constants.ContextIdentityKey, context_helper.Identity{
			UserID:   claims.UserID,
			Username: claims.Username,
		})

		c.Next()
	}
}
