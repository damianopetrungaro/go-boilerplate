package db

import (
	"database/sql"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq" // Postgres driver
)

var url = os.Getenv("DATABASE_URL")
var driver = os.Getenv("DATABASE_DRIVER")
var maxIdleConn = os.Getenv("DATABASE_MAX_IDLE_CONN")
var maxOpenConn = os.Getenv("DATABASE_MAX_OPEN_CONN")
var maxConnLifetime = os.Getenv("DATABASE_MAX_CONN_LIFETIME")

// New Create a database connection using the environment variable to define the database driver and url
// it returns an error when an error occurs establishing the connection
func New() (*sql.DB, error) {
	db, err := sql.Open(driver, url)
	if err != nil {
		return nil, err
	}

	if i, err := strconv.Atoi(maxIdleConn); err != nil {
		db.SetMaxIdleConns(i)
	}

	if i, err := strconv.Atoi(maxOpenConn); err != nil {
		db.SetMaxIdleConns(i)
	}

	if i, err := strconv.Atoi(maxConnLifetime); err != nil {
		db.SetConnMaxLifetime(time.Duration(i) * time.Minute)
	}

	return db, nil
}
