package respwr

import (
	"encoding/json"
	"net/http"
)

type httpResponseWriter struct{}

type HTTPResponseWriter interface {
	// WriteResponseData write success response contain data field
	WriteResponseData(w http.ResponseWriter, code int, data any)
	// WriteResponseData write fail response contain error message
	WriteErrMessage(w http.ResponseWriter, code int, message string)
	// WriteResponseData write fail response contain error message and error details
	WriteErrDetails(w http.ResponseWriter, code int, message string, errMap map[string]string)
}

type ResponseData[T any] struct {
	Data T `json:"data"`
}

type ErrMessage struct {
	Message string `json:"message"`
}

type ErrDetails struct {
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"`
}

func NewHTTPResponseWriter() HTTPResponseWriter {
	return &httpResponseWriter{}
}

func (rw *httpResponseWriter) setHeader(w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
}

func (rw *httpResponseWriter) handleEncodingErr(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal server error"))
}

func (rw *httpResponseWriter) WriteResponseData(w http.ResponseWriter, code int, data any) {
	rw.setHeader(w, code)
	if err := json.NewEncoder(w).Encode(ResponseData[any]{
		Data: data,
	}); err != nil {
		rw.handleEncodingErr(w)
	}
}

func (rw *httpResponseWriter) WriteErrMessage(w http.ResponseWriter, code int, message string) {
	rw.setHeader(w, code)
	if err := json.NewEncoder(w).Encode(ErrMessage{
		Message: message,
	}); err != nil {
		rw.handleEncodingErr(w)
	}
}

func (rw *httpResponseWriter) WriteErrDetails(w http.ResponseWriter, code int, message string, errMap map[string]string) {
	rw.setHeader(w, code)
	if err := json.NewEncoder(w).Encode(ErrDetails{
		Message: message,
		Errors:  errMap,
	}); err != nil {
		rw.handleEncodingErr(w)
	}
}
