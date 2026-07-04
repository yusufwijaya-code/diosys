package context_helper

import (
	"portfolio-api/base/helpers/error_helper"
	"portfolio-api/constants"

	"github.com/gin-gonic/gin"
)

// Identity holds the authenticated user information extracted from the JWT.
type Identity struct {
	UserID   int    `json:"userID"`
	Username string `json:"username"`
}

// GetIdentity reads the authenticated identity placed in the context by the
// JWT middleware.
func GetIdentity(c *gin.Context) (Identity, error) {
	value, exists := c.Get(constants.ContextIdentityKey)
	if !exists {
		return Identity{}, error_helper.Unauthorized("missing authentication context")
	}

	identity, ok := value.(Identity)
	if !ok {
		return Identity{}, error_helper.Unauthorized("invalid authentication context")
	}

	return identity, nil
}
