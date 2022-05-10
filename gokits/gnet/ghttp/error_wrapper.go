package ghttp

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse container status code and error code message,
// it will encode to a json string and add to http body
type ErrorResponse struct {
	statusCode   int                    `json:"-"`
	ErrorCode    string                 `json:"error_code"`
	ErrorMessage string                 `json:"error_message"`
	ErrorParams  map[string]interface{} `json:"error_params"`
}

// Error implements error interface
func (er *ErrorResponse) Error() string {
	bs, _ := json.Marshal(er)
	return string(bs)
}

// WithMsg add error message
func (er *ErrorResponse) WithMsg(errMsg string) *ErrorResponse {
	er.ErrorMessage = errMsg
	return er
}

// WithParam add error parameters with key and value
func (er *ErrorResponse) WithParam(key string, value interface{}) *ErrorResponse {
	er.ErrorParams[key] = value
	return er
}

// NewErrorRes creates a instance of ErrorResponse
func NewErrorRes(statusCode int) *ErrorResponse {
	return &ErrorResponse{statusCode: statusCode, ErrorParams: make(map[string]interface{})}
}

// HandlerFunc is http handler function return error
type HandlerFunc func(http.ResponseWriter, *http.Request) error

func WrapperHandler(handlerFunc HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handlerFunc(w, r)
		if err == nil {
			return
		}

		if errRsp, ok := err.(*ErrorResponse); ok {
			w.WriteHeader(errRsp.statusCode)
			if err := json.NewEncoder(w).Encode(errRsp); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
	}
}
