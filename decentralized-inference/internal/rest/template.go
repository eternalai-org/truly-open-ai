package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	StatusSuccess int32 = 1
	StatusFailed  int32 = -1

	LocalStatusInBody string = "response_status_in_body"
)

type Err struct {
	err          error
	HttpCode     int    `json:"-"`
	ErrorMessage string `json:"error_message"`
}

func (e *Err) Error() string {
	return e.err.Error()
}

func NewHttpErr(err error, httpCode int, message ...string) *Err {
	msg := ""
	if len(message) == 1 {
		msg = message[0]
	}
	return &Err{
		err:          err,
		ErrorMessage: msg,
		HttpCode:     httpCode,
	}
}

type Response struct {
	HttpCode int         `json:"-"`
	Message  string      `json:"message,omitempty"`
	Error    string      `json:"error,omitempty"`
	Status   int32       `json:"status"`
	Data     interface{} `json:"data,omitempty"`
}

type StreamResponse struct {
	IsNotStream bool
	Data        interface{}
}

type handlerFunc func(c *gin.Context) (interface{}, error)

func ResponseJSON(h handlerFunc, c *gin.Context) {
	resp, err := h(c)
	if err != nil {
		errorHandler(c, err)
		return
	}
	c.JSON(http.StatusOK, NewSuccessResponse(resp))
}

func StreamResponseJSON(h handlerFunc, c *gin.Context) {
	resp, err := h(c)
	if err != nil {
		errorHandler(c, err)
		return
	}
	_resp := resp.(*StreamResponse)

	if _resp.IsNotStream {
		c.JSON(http.StatusOK, NewSuccessResponse(_resp.Data))
		return
	}

}

func errorHandler(c *gin.Context, err error) {
	resp := FromErr(c, err)
	if resp == nil {
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	errCode := http.StatusInternalServerError
	if resp.HttpCode != 0 {
		errCode = resp.HttpCode
	}
	c.JSON(errCode, resp)
}

func NewSuccessResponse(data interface{}) *Response {
	return &Response{
		Status: StatusSuccess,
		Data:   data,
	}
}

func FromErr(c *gin.Context, err error) *Response {
	er := &Response{
		Status: StatusFailed,
	}

	if e, ok := err.(*Err); ok {
		er.HttpCode = e.HttpCode
		er.SetFields(er.WithMessage(e.ErrorMessage))
		er.SetFields(er.WithError(e.Error()))
		return er
	}

	er.SetFields(er.WithMessage(err.Error()))
	er.SetFields(er.WithError(err.Error()))
	return er
}

func (e *Response) SetFields(fns ...func(*Response)) {
	for _, fn := range fns {
		fn(e)
	}
}

func (e *Response) WithError(err string) func(*Response) {
	return func(r *Response) {
		r.Error = err
	}
}

func (e *Response) WithMessage(message string) func(*Response) {
	return func(r *Response) {
		r.Message = message
	}
}
