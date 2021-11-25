package http

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	"github.com/masterraf21/dss-backend/utils/errors"
	logger "github.com/sirupsen/logrus"
)

type (
	// Response - structure response
	Response struct {
		Ctx          echo.Context `json:"-"`
		RequestID    string       `json:"request_id,omitempty"`
		Code         int          `json:"code"`
		Message      string       `json:"message"`
		ErrorMessage string       `json:"error_message,omitempty"`
		Data         interface{}  `json:"data,omitempty"`
		// App          string       `json:"app",omit`
	}
)

// NewResponse will create a response
func NewResponse(code int, data interface{}) *Response {
	return &Response{
		Code: code,
		// Message: msg,
		Data: data,
	}
}

func NewError(code int) *Response {
	return &Response{
		Code: code,
	}
}

// WriteResponse will actually write the response
func (r *Response) WriteResponse(
	ctx echo.Context,
) (err error) {
	// var (
	// 	request                                 = ctx.Request()
	// 	endpointURL                             = request.RequestURI
	// 	requestMethod                           = request.Method
	// 	requestPayloadByte, responsePayloadByte []byte
	// 	requestPayloadStr, responsePayloadStr   string
	// 	authType                                string
	// )

	ctx.Response().Header().Add("Access-Control-Expose-Headers", "Content-Length, Content-Range, Date, date")
	ctx.Response().Header().Add("Content-Type", "application/json")
	ctx.Response().WriteHeader(r.Code)
	return json.NewEncoder(ctx.Response()).Encode(r)
}

func (r *Response) WriteError(
	ctx echo.Context,
	message string,
	err error,
) error {
	switch r.Code {
	case http.StatusBadRequest:
		logger.Error(errors.BadRequestError(ctx.Request().RequestURI, ctx.Request().Method, message+", "+err.Error()))
	case http.StatusUnprocessableEntity:
		logger.Error(errors.UnprocessableError(ctx.Request().RequestURI, ctx.Request().Method, message+", "+err.Error()))
	case http.StatusNotFound:
		logger.Error(errors.NotFoundError(ctx.Request().RequestURI, ctx.Request().Method, message+", "+err.Error()))
	case http.StatusInternalServerError:
		logger.Error(errors.DefaultError(ctx.Request().RequestURI, ctx.Request().Method, message+", "+err.Error()))
	default:
		logger.Error(errors.DefaultError(ctx.Request().RequestURI, ctx.Request().Method, message+", "+err.Error()))
	}

	ctx.Response().Header().Add("Access-Control-Expose-Headers", "Content-Length, Content-Range, Date, date")
	ctx.Response().Header().Add("Content-Type", "application/json")
	ctx.Response().WriteHeader(r.Code)
	return json.NewEncoder(ctx.Response()).Encode(r)
}
