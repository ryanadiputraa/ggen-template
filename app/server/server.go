package server

import (
	"database/sql"
	"net/http"

	"github.com/ryanadiputraa/ggen-template/config"
	"github.com/ryanadiputraa/ggen-template/pkg/logger"
	"github.com/ryanadiputraa/ggen-template/pkg/middleware"
	"github.com/ryanadiputraa/ggen-template/pkg/respwr"
)

type Server struct {
	config config.Config
	log    logger.Logger
	web    *http.ServeMux
	db     *sql.DB
	respwr respwr.HTTPResponseWriter
}

func NewServer(config config.Config, logger logger.Logger, db *sql.DB) *http.Server {
	s := Server{
		config: config,
		log:    logger,
		web:    http.NewServeMux(),
		db:     db,
		respwr: respwr.NewHTTPResponseWriter(),
	}

	s.setupHandlers()
	handler := s.Use(
		middleware.CORSMiddleware,
		middleware.TimeoutMiddleware,
		middleware.ThrottleMiddleware,
	)

	return &http.Server{
		Addr:    config.Port,
		Handler: handler,
	}
}

// Use return http.Handler with attached middlewares
func (s *Server) Use(middlewares ...func(handler http.Handler) http.Handler) (handler http.Handler) {
	handler = s.web
	for _, m := range middlewares {
		handler = m(handler)
	}
	return
}
