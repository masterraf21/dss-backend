package http

import (
	"encoding/json"

	"github.com/labstack/echo"
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
		App          string       `json:"app"`
	}
)

// NewResponse will create a response
func NewResponse(code int, msg string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: msg,
		Data:    data,
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
	ctx.Response().WriteHeader(r.Code)
	return json.NewEncoder(ctx.Response()).Encode(r)
}
