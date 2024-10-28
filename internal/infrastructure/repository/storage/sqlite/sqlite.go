package sqlite

import (
	"database/sql"
	"errors"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/alexanderiand/notification-service/pkg/config"
)

const (
	// Retry
	waitingTime = time.Second * 1
	attempts    = 3
)

var (
	ErrNilStructPointer = errors.New("error, nil struct pointer")
	ErrInvalidDBPath    = errors.New("error, invalid database file path")
)

type SQLite struct {
	DB *sql.DB
}

// New SQLite client constructor, return *SQLite, error
// If db file path is invalid, return ErrInvalidDBPath
func New(cfg *config.Config) (*SQLite, error) {
	if cfg == nil {
		return nil, ErrNilStructPointer
	}
	db := &SQLite{}

	println(cfg.DatabaseFilePath)

	err := DoWithTries(func() error {
		sqlite, err := sql.Open("sqlite3", cfg.DatabaseFilePath)
		if err != nil {
			return err
		}
		db.DB = sqlite

		return nil
	}, attempts, waitingTime)

	if err != nil {
		return nil, err
	}

	return db, nil
}

// DoWithTries implement Retry pattern
func DoWithTries(fn func() error, attempts int, delay time.Duration) error {
	for attempts > 0 {
		if err := fn(); err != nil {
			time.Sleep(delay)
			attempts--
			continue
		}
		return nil
	}
	return errors.New("error, 0 connection attempts left: the database is not connected")
}
