package httperror

import (
	"fmt"
	"net/http"
)

type HTTPError struct {
	HTTPCode int
	ID       string
	Message  string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("[%d] [%s] %s", e.HTTPCode, e.ID, e.Message)
}

func Default() *HTTPError {
	return &HTTPError{
		HTTPCode: http.StatusInternalServerError,
		ID:       "Default",
		Message:  "internal server error",
	}
}

func InvalidToken() *HTTPError {
	return &HTTPError{
		HTTPCode: http.StatusUnauthorized,
		ID:       "InvalidToken",
		Message:  "invalid token",
	}
}
