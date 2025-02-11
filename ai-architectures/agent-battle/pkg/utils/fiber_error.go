package utils

import (
	"encoding/json"
	"fmt"

	"agent-battle/pkg/constants"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Err struct {
	err          error
	SystemError  string `json:"error"`
	ErrorCode    int32  `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	HttpCode     int    `json:"-"`
}

var ErrResourceConflict = NewHttpErr(nil, constants.ErrPreconditionFailed, fiber.StatusPreconditionFailed, "Resource has been modified by other request")

type aliasError Err

func (e *Err) MarshalJSON() ([]byte, error) {
	err := ""
	if e.err != nil {
		err = e.err.Error()
	}
	return json.Marshal(&struct {
		Error string `json:"error,omitempty"`
		*aliasError
	}{
		Error:      err,
		aliasError: (*aliasError)(e),
	})
}

func NewBadRequestErr(code int32, message string, err ...error) *Err {
	var e error
	if len(err) == 1 {
		e = err[0]
	}
	return &Err{
		err:          e,
		ErrorCode:    code,
		ErrorMessage: message,
		HttpCode:     fiber.StatusBadRequest,
	}
}

func NewErr(err error, code int32, message ...string) *Err {
	msg := ""
	if len(message) == 1 {
		msg = message[0]
	}
	return &Err{
		err:          err,
		ErrorCode:    code,
		ErrorMessage: msg,
		HttpCode:     fiber.StatusInternalServerError,
	}
}

func NewHttpErr(err error, code int32, httpCode int, message ...string) *Err {
	msg := ""
	if len(message) == 1 {
		msg = message[0]
	}
	return &Err{
		err:          err,
		ErrorCode:    code,
		ErrorMessage: msg,
		HttpCode:     httpCode,
	}
}

func (e *Err) Error() string {
	return fmt.Sprintf("%v ,code=%d, message=%s", e.err, e.ErrorCode, e.ErrorMessage)
}

func FromErr(err error) *Err {
	if e, ok := err.(*Err); ok {
		return e
	}

	er := &Err{
		err:         err,
		ErrorCode:   50_000,
		SystemError: err.Error(),
		HttpCode:    fiber.StatusInternalServerError,
	}

	if e, ok := err.(*fiber.Error); ok {
		er.HttpCode = e.Code
		er.ErrorMessage = e.Message
		return er
	}

	if s, ok := status.FromError(err); ok {
		switch s.Code() {
		case codes.NotFound:
			er.HttpCode = fiber.StatusNotFound
		case codes.InvalidArgument:
			er.HttpCode = fiber.StatusBadRequest
		case codes.DeadlineExceeded:
			er.HttpCode = fiber.StatusRequestTimeout
		case codes.FailedPrecondition:
			er.HttpCode = fiber.StatusPreconditionFailed
		case codes.AlreadyExists:
			er.HttpCode = fiber.StatusConflict
		}
		er.ErrorMessage = s.Message()
	}

	if er.ErrorMessage == "" {
		er.ErrorMessage = er.SystemError
	}
	return er
}
