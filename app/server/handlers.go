package server

import (
	healthcheckHandler "github.com/ryanadiputraa/ggen-template/app/healthcheck/handler"
	"github.com/ryanadiputraa/ggen-template/pkg/respwr"
)

func (s *Server) setupHandlers() {
	respwr := respwr.NewHTTPResponseWriter()
	healthcheckHandler := healthcheckHandler.NewHTTPHandler(respwr)
	s.web.Handle("GET /healthcheck", healthcheckHandler.Healthcheck())
}
