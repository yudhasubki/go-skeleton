package libs

import "reflect"

type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ApiResponse(code, message string, data interface{}) *Response {
	resp := Response{}
	resp.Code = code
	resp.Message = message
	if data != nil || reflect.ValueOf(data).Kind() == reflect.Ptr {
		resp.Data = data
	}
	return &resp
}
