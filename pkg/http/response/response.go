package response

import (
	"errors"

	"github.com/Hidayathamir/goproj/pkg/http/httperror"
	"github.com/gin-gonic/gin"
)

type ResHTTP[data any] struct {
	Data     data   `json:"data"`
	ErrorID  string `json:"error_id"`
	ErrorMsg string `json:"error_message"`
}

func Data(c *gin.Context, status int, data any) {
	res := ResHTTP[any]{}
	res.Data = data
	c.JSON(status, res)
}

func Error(c *gin.Context, err error) {
	if err == nil {
		return
	}

	httpErr := LoadErrAsHTTPError(err)

	res := ResHTTP[any]{}
	res.ErrorID = httpErr.ID
	res.ErrorMsg = httpErr.Message

	c.JSON(httpErr.HTTPCode, res)
}

// LoadErrAsHTTPError load error as HTTPError with use Default().
func LoadErrAsHTTPError(err error) *httperror.HTTPError {
	httpErr := httperror.Default()
	errors.As(err, &httpErr)

	return httpErr
}
