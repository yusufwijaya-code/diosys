package auth_handler

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"net/url"
	"strings"

	"portfolio-api/base/helpers/context_helper"
	"portfolio-api/base/helpers/http_helper"
	"portfolio-api/config"
	"portfolio-api/constants"
	"portfolio-api/modules/auth/auth_service"
	"portfolio-api/modules/user/user_service"

	"github.com/gin-gonic/gin"
)

const oauthStateCookie = "diosys_oauth_state"

// AuthHandler exposes authentication endpoints.
type AuthHandler interface {
	GoogleRedirect(c *gin.Context)
	GoogleCallback(c *gin.Context)
	Me(c *gin.Context)
}

type authHandlerImpl struct {
	authService auth_service.AuthService
	userService user_service.UserService
	config      config.AppConfig
}

// NewAuthHandler builds an AuthHandler.
func NewAuthHandler(authService auth_service.AuthService, userService user_service.UserService, cfg config.AppConfig) AuthHandler {
	return &authHandlerImpl{authService: authService, userService: userService, config: cfg}
}

// GoogleRedirect starts the OAuth flow by sending the browser to Google.
func (h *authHandlerImpl) GoogleRedirect(c *gin.Context) {
	state := randomState()
	// Short-lived state cookie for CSRF protection.
	c.SetCookie(oauthStateCookie, state, 300, "/", "", false, true)
	c.Redirect(http.StatusTemporaryRedirect, h.authService.BuildAuthURL(state))
}

// GoogleCallback handles Google's redirect back, then bounces to the frontend
// login page with either the issued token or an error message.
func (h *authHandlerImpl) GoogleCallback(c *gin.Context) {
	loginURL := strings.TrimRight(h.config.AppFrontendUrl, "/") + "/admin/login"

	if errParam := c.Query("error"); errParam != "" {
		h.redirectWithError(c, loginURL, "Google sign-in was cancelled.")
		return
	}

	expectedState, _ := c.Cookie(oauthStateCookie)
	c.SetCookie(oauthStateCookie, "", -1, "/", "", false, true)
	if expectedState == "" || c.Query("state") != expectedState {
		h.redirectWithError(c, loginURL, "Invalid sign-in state. Please try again.")
		return
	}

	code := c.Query("code")
	if code == "" {
		h.redirectWithError(c, loginURL, "Missing authorization code.")
		return
	}

	response, err := h.authService.HandleCallback(c.Request.Context(), code)
	if err != nil {
		h.redirectWithError(c, loginURL, err.Error())
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, loginURL+"?token="+url.QueryEscape(response.AccessToken))
}

func (h *authHandlerImpl) Me(c *gin.Context) {
	identity, err := context_helper.GetIdentity(c)
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}

	user, err := h.userService.GetByID(identity.UserID)
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}

	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "OK", user)
}

func (h *authHandlerImpl) redirectWithError(c *gin.Context, loginURL, message string) {
	c.Redirect(http.StatusTemporaryRedirect, loginURL+"?error="+url.QueryEscape(message))
}

func randomState() string {
	buf := make([]byte, 16)
	if _, err := rand.Read(buf); err != nil {
		return "diosys-state"
	}
	return hex.EncodeToString(buf)
}
