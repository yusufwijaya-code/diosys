package constants

// Response codes returned in the API envelope.
const (
	EC_SUCCESS          = "200"
	EC_CREATED          = "201"
	EC_VALIDATION_ERROR = "400"
	EC_UNAUTHORIZED     = "401"
	EC_FORBIDDEN        = "403"
	EC_NOT_FOUND        = "404"
	EC_INTERNAL_ERROR   = "500"
)

// ContextIdentityKey is the gin context key holding the authenticated identity.
const ContextIdentityKey = "identity"
