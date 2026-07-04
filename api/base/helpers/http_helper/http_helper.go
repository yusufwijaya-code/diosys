package http_helper

import (
	"errors"
	"net/http"
	"time"

	"portfolio-api/base/helpers/error_helper"
	"portfolio-api/constants"

	"github.com/gin-gonic/gin"
)

// Response is the standard API envelope returned by every endpoint.
type Response struct {
	Path      string                    `json:"path"`
	Timestamp string                    `json:"timestamp"`
	Status    string                    `json:"status"`
	Code      string                    `json:"code"`
	Message   string                    `json:"message"`
	Result    interface{}               `json:"result"`
	Errors    []error_helper.FieldError `json:"errors"`
}

// PaginationResult wraps a paginated list together with its metadata.
type PaginationResult[T any] struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Count int `json:"count"`
	Data  []T `json:"data"`
}

func newResponse(c *gin.Context, status, code, message string, result interface{}, fieldErrors []error_helper.FieldError) Response {
	return Response{
		Path:      c.Request.RequestURI,
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		Status:    status,
		Code:      code,
		Message:   message,
		Result:    result,
		Errors:    fieldErrors,
	}
}

// SuccessResponse writes a 200 success envelope.
func SuccessResponse(c *gin.Context, code, message string, result interface{}) {
	if message == "" {
		message = "OK"
	}
	c.JSON(http.StatusOK, newResponse(c, "ok", code, message, result, nil))
	c.Abort()
}

// SuccessPaginationResponse writes a 200 success envelope with pagination metadata.
func SuccessPaginationResponse[T any](c *gin.Context, code, message string, data []T, count, page, limit int) {
	if message == "" {
		message = "OK"
	}
	if data == nil {
		data = []T{}
	}
	result := PaginationResult[T]{Page: page, Limit: limit, Count: count, Data: data}
	c.JSON(http.StatusOK, newResponse(c, "ok", code, message, result, nil))
	c.Abort()
}

// HttpErrorResponse maps an error into the proper HTTP envelope. Known AppError
// values keep their status/code; anything else becomes a 500.
func HttpErrorResponse(c *gin.Context, err error) {
	var appErr *error_helper.AppError
	if errors.As(err, &appErr) {
		c.JSON(appErr.HttpStatus, newResponse(c, "fail", appErr.Code, appErr.Message, nil, appErr.Fields))
		c.Abort()
		return
	}

	c.JSON(
		http.StatusInternalServerError,
		newResponse(c, "fail", constants.EC_INTERNAL_ERROR, err.Error(), nil, nil),
	)
	c.Abort()
}
