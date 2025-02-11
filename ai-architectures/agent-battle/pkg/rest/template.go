package rest

import (

"agent-battle/internal/core/middleware"
"errors"

	"github.com/gofiber/fiber/v2"
)

type handlerFunc func(ctx *fiber.Ctx, currentUserEmail string) (interface{}, error)

type fiberHandlerTemplate struct {
	handlerFunc handlerFunc
}

const (
	StatusSuccess int32 = 1
	StatusFailed  int32 = -1
)

type Response struct {
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
	Status  int32       `json:"status"`
	Data    interface{} `json:"data,omitempty"`
}

func NewFiberHandlerTemplate(handlerFunc handlerFunc) *fiberHandlerTemplate {
	return &fiberHandlerTemplate{
		handlerFunc: handlerFunc,
	}
}

func (s *fiberHandlerTemplate) ResponseJSON(ctx *fiber.Ctx) error {
	resp, err := s.handlerFunc(ctx, middleware.GetCurrentUserEmail(ctx))
	if err != nil {
		return err
	}
	return ctx.JSON(NewSuccessResponse(resp))
}

func NewSuccessResponse(data interface{}) *Response {
	return &Response{
		Status: StatusSuccess,
		Data:   data,
	}
}

func (s *fiberHandlerTemplate) SetEmptyResponse(ctx *fiber.Ctx) error {
	_, err := s.handlerFunc(ctx, middleware.GetCurrentUserEmail(ctx))
	if err != nil {
		return err
	}
	ctx.Response().SetStatusCode(fiber.StatusNoContent)
	return nil
}

func (s *fiberHandlerTemplate) SendString(ctx *fiber.Ctx) error {
	resp, err := s.handlerFunc(ctx, middleware.GetCurrentUserEmail(ctx))
	if err != nil {
		return err
	}
	val, ok := resp.(string)
	if !ok {
		return errors.New("Response must is string")
	}
	return ctx.SendString(val)
}

func (s *fiberHandlerTemplate) Send(ctx *fiber.Ctx) error {
	resp, err := s.handlerFunc(ctx, middleware.GetCurrentUserEmail(ctx))
	if err != nil {
		return err
	}
	val, ok := resp.([]byte)
	if !ok {
		return errors.New("Response must is byte")
	}
	return ctx.Send(val)
}
