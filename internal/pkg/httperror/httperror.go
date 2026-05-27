package httperror

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/behnamdehghannejad/vendorservice/internal/pkg/apperror"
	"github.com/gin-gonic/gin"
)

func Handle(err error) (gin.H, int) {
	return getErrorBody(err), getStatus(err)
}

func getStatus(err error) int {
	var appErr *apperror.AppError

	var unmarshallTypeErr *json.UnmarshalTypeError
	if errors.As(err, &unmarshallTypeErr) {
		return http.StatusBadRequest
	}

	if !errors.As(err, &appErr) {
		return http.StatusInternalServerError
	}
	errType := appErr.GetErrorType()

	return mapAppErrorTypeToHttpStatus(errType)
}

func mapAppErrorTypeToHttpStatus(errType apperror.ErrorType) int {
	switch errType {
	case apperror.BadRequest:
		return http.StatusBadRequest
	case apperror.Forbidden:
		return http.StatusForbidden
	case apperror.UnExpected:
		return http.StatusInternalServerError
	case apperror.NotFound:
		return http.StatusNotFound
	case apperror.Duplicate:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func getErrorBody(err error) gin.H {
	code := getStatus(err)
	switch code {
	case http.StatusBadRequest:
		return gin.H{
			"error": "input is wrong",
		}
	case http.StatusNotFound:
		return gin.H{
			"error": "no record found",
		}
	case http.StatusConflict:
		return gin.H{
			"error": "this record exist",
		}
	}
	return gin.H{
		"error": "something went wrong",
	}
}
