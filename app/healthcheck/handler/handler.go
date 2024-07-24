package handler

import (
	"net/http"

	"github.com/ryanadiputraa/ggen-template/app/healthcheck"
	"github.com/ryanadiputraa/ggen-template/pkg/respwr"
)

type handler struct {
	respwr respwr.HTTPResponseWriter
}

func NewHTTPHandler(respwr respwr.HTTPResponseWriter) handler {
	return handler{
		respwr: respwr,
	}
}

func (h *handler) Healthcheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := healthcheck.Healthcheck{
			Status: "ok",
		}
		h.respwr.WriteResponseData(w, http.StatusOK, data)
	}
}
