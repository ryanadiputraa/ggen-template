package test

import (
	"database/sql"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ryanadiputraa/ggen-template/app/server"
	"github.com/ryanadiputraa/ggen-template/config"
	"github.com/ryanadiputraa/ggen-template/pkg/logger"
)

func newServer(db *sql.DB) *http.Server {
	logger := logger.New(time.UTC, os.Stderr)
	config := config.Config{
		Port: ":8080",
	}

	return server.NewServer(config, logger, db)
}

func newMockDB(t *testing.T) (db *sql.DB, mock sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("fail to create mock db conn. Err: %v", err.Error())
	}
	return
}

func runServer(server *http.Server, srvErr chan error) {
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			srvErr <- err
		}
	}()
}
