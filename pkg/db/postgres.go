package db

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

const (
	maxOpenConns    = 60
	connMaxLifeTime = 120
	maxIdleConn     = 30
	connMaxIdleTime = 20
)

func NewPostgres(dsn string) (db *sql.DB, err error) {
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		return
	}

	// Health check by pinging database with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err = db.PingContext(ctx); err != nil {
		return
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(connMaxLifeTime)
	db.SetMaxIdleConns(maxIdleConn)
	db.SetConnMaxIdleTime(connMaxIdleTime)
	return
}
