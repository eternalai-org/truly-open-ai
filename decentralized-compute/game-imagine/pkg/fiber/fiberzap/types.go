package fiberzap

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap/zapcore"
)

func getAllowedHeaders() map[string]bool {
	return map[string]bool{
		"User-Agent":      true,
		"Content-Type":    true,
		"Origin":          true,
		"Referer":         true,
		"Accept-Language": true,
		"Authorization":   true,
	}
}

type resp struct {
	code       int
	contenType string
	body       string
}

func Resp(r *fasthttp.Response) *resp {
	var (
		contenType = bytes.NewBuffer(r.Header.ContentType()).String()
		body       []byte
	)
	if strings.Contains(contenType, fiber.MIMEApplicationJSON) {
		buffer := new(bytes.Buffer)
		err := json.Compact(buffer, r.Body())
		if err != nil {
			body, _ = fasthttp.AppendUnbrotliBytes(body, r.Body())
		} else {
			body = buffer.Bytes()
		}
	}
	rs := &resp{
		code:       r.StatusCode(),
		contenType: contenType,
		body:       bytes.NewBuffer(body).String(),
	}
	return rs
}

func (r *resp) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("type", r.contenType)
	enc.AddInt("code", r.code)
	if r.body != "" {
		enc.AddString("body", r.body)
	}
	return nil
}

type req struct {
	body    string
	query   string
	user    string
	ip      string
	method  string
	path    string
	route   string
	host    string
	scheme  string
	referer string
	headers *headerbag
}

func Req(c *fiber.Ctx) *req {
	reqq := c.Request()
	var body []byte
	buffer := new(bytes.Buffer)
	err := json.Compact(buffer, reqq.Body())
	if err != nil {
		body = reqq.Body()
	} else {
		body = buffer.Bytes()
	}

	headers := &headerbag{
		vals: make(map[string]string),
	}
	allowedHeaders := getAllowedHeaders()
	reqq.Header.VisitAll(func(key, val []byte) {
		k := bytes.NewBuffer(key).String()
		if _, exist := allowedHeaders[k]; exist {
			headers.vals[strings.ToLower(k)] = bytes.NewBuffer(val).String()
		}
	})

	var userEmail string
	if u := c.Locals("username"); u != nil {
		userEmail = u.(string)
	}

	return &req{
		body:    bytes.NewBuffer(body).String(),
		query:   bytes.NewBuffer(reqq.URI().QueryString()).String(),
		host:    bytes.NewBuffer(reqq.URI().Host()).String(),
		scheme:  bytes.NewBuffer(reqq.URI().Scheme()).String(),
		path:    bytes.NewBuffer(reqq.URI().Path()).String(),
		referer: bytes.NewBuffer(reqq.Header.Referer()).String(),
		route:   c.Route().Path,
		headers: headers,
		ip:      c.IP(),
		method:  c.Method(),
		user:    userEmail,
	}
}

func (r *req) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("url_details.host", r.host)
	enc.AddString("url_details.path", r.path)
	enc.AddString("url_details.route", r.route)
	enc.AddString("url_details.query", r.query)
	enc.AddString("url_details.scheme", r.scheme)
	enc.AddInt("url_details.port", 80)
	enc.AddString("ip", r.ip)
	enc.AddString("method", r.method)
	enc.AddString("referer", r.referer)

	if r.body != "" {
		enc.AddString("body", r.body)
	}

	if r.user != "" {
		enc.AddString("user", r.user)
	}

	err := enc.AddObject("headers", r.headers)
	if err != nil {
		return err
	}

	return nil
}

type headerbag struct {
	vals map[string]string
}

func (h *headerbag) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	for k, v := range h.vals {
		enc.AddString(k, v)
	}

	return nil
}
