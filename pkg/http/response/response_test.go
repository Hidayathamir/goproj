package response_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/Hidayathamir/goproj/pkg/http/httperror"
	"github.com/Hidayathamir/goproj/pkg/http/response"
	"github.com/stretchr/testify/assert"
)

func TestLoadErrAsHTTPError_1(t *testing.T) {
	var err error
	err = errors.New("dummy err 1")
	err = fmt.Errorf("wrap: %w", err)
	err = errors.Join(httperror.InvalidToken(), err)
	err = fmt.Errorf("wrap2: %w", err)

	httpErr := response.LoadErrAsHTTPError(err)

	assert.Equal(t, httperror.InvalidToken().ID, httpErr.ID)
	assert.Equal(t, httperror.InvalidToken().Message, httpErr.Message)
}

func TestLoadErrAsHTTPError_2(t *testing.T) {
	var err error
	err = errors.New("dummy err 1")
	err = fmt.Errorf("wrap: %w", err)
	err = fmt.Errorf("wrap2: %w", err)

	httpErr := response.LoadErrAsHTTPError(err)

	assert.Equal(t, httperror.Default().ID, httpErr.ID)
	assert.Equal(t, httperror.Default().Message, httpErr.Message)
}

func TestLoadErrAsHTTPError_3(t *testing.T) {
	var err error
	err = errors.New("dummy err 1")
	err = fmt.Errorf("wrap: %w", err)
	err = errors.Join(httperror.InvalidToken(), err)
	err = fmt.Errorf("wrap2: %w", err)
	err = errors.Join(httperror.Default(), err)

	httpErr := response.LoadErrAsHTTPError(err)

	assert.Equal(t, httperror.Default().ID, httpErr.ID)
	assert.Equal(t, httperror.Default().Message, httpErr.Message)
}

func TestLoadErrAsHTTPError_4(t *testing.T) {
	var err error
	err = errors.New("dummy err 1")
	err = fmt.Errorf("wrap: %w", err)
	err = errors.Join(httperror.InvalidToken(), err)
	err = fmt.Errorf("wrap2: %w", err)
	err = errors.Join(httperror.Default(), err)
	err = errors.Join(httperror.InvalidToken(), err)

	httpErr := response.LoadErrAsHTTPError(err)

	assert.Equal(t, httperror.InvalidToken().ID, httpErr.ID)
	assert.Equal(t, httperror.InvalidToken().Message, httpErr.Message)
}

func TestLoadErrAsHTTPError_5(t *testing.T) {
	var err error = httperror.InvalidToken()

	httpErr := response.LoadErrAsHTTPError(err)

	assert.Equal(t, httperror.InvalidToken().ID, httpErr.ID)
	assert.Equal(t, httperror.InvalidToken().Message, httpErr.Message)
}

func TestLoadErrAsHTTPError_6(t *testing.T) {
	var err error

	httpErr := response.LoadErrAsHTTPError(err)

	assert.Equal(t, httperror.Default().ID, httpErr.ID)
	assert.Equal(t, httperror.Default().Message, httpErr.Message)
}

func TestLoadErrAsHTTPError_7(t *testing.T) {
	var err error
	err = httperror.InvalidToken()
	err = fmt.Errorf("wrap1: %w", err)

	httpErr := response.LoadErrAsHTTPError(err)

	assert.Equal(t, httperror.InvalidToken().ID, httpErr.ID)
	assert.Equal(t, httperror.InvalidToken().Message, httpErr.Message)
}
