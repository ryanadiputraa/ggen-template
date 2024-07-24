package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ryanadiputraa/ggen-template/app/server"
	"github.com/ryanadiputraa/ggen-template/config"
	"github.com/ryanadiputraa/ggen-template/pkg/db"
	"github.com/ryanadiputraa/ggen-template/pkg/logger"
)

func Run() {
	logger := logger.New(time.UTC, os.Stderr)

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("error loading config file. Err:", err)
	}

	db, err := db.NewPostgres(config.PostgresDSN)
	if err != nil {
		logger.Fatal("error opening postgres connection. Err:", err)
	}

	s := server.NewServer(config, logger, db)

	go func() {
		done := make(chan os.Signal, 1)
		signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-done
		logger.Info("start shutdown")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		if err := s.Shutdown(ctx); err == context.DeadlineExceeded {
			logger.Fatal("error shuting down server: context exeeded")
		}
	}()

	logger.Info("starting server on port", config.Port)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal("error statrting server:", err)
	}
}
