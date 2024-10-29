package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/alexanderiand/notification-service/internal/entity"
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

// Storage
type Storage interface {
	NotificationProvider
	NotificationSaver
}

type NotificationProvider interface {
	GetAllEvents() (events []*entity.Event, err error)
}

type NotificationSaver interface {
	SaveEvent(event *entity.Event) (evenID int, err error)
}

type SQLite struct {
	Storage
	*sql.DB
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

// SaveEvent into events table of the sqlite3 database, return  the saved event.id, nil
// If something going wrong, return 0, and error
func (s *SQLite) SaveEvent(e *entity.Event) (evenID int, err error) {
	op := "internal.infrastructure.storage.sqlite.SaveEvent"

	stmt, err := s.Prepare("INSERT INTO events(order_type, session_id, card, event_date, website_url) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return 0, fmt.Errorf("%s %w", op, err)
	}

	res, err := stmt.Exec(e.OrderType, e.SessionID, e.Card, e.EventDate, e.WebSiteURL)
	if err != nil {
		return 0, fmt.Errorf("%s %w", op, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s %w", op, err)
	}

	return int(id), nil
}

// GetAllEvents select all events from events table of db, return []*entity.Events, nil
// If something going wrong, return nil, and error
func (s *SQLite) GetAllEvents() (events []*entity.Event, err error) {
	op := "internal.infrastructure.storage.sqlite.GetAllEvents"

	stmt, err := s.Query("SELECT id, order_type, session_id, card, event_date, website_url, created_at FROM events")
	if err != nil {
		return nil, fmt.Errorf("%s %w", op, err)
	}

	for stmt.Next() {
		e := &entity.Event{}
		err := stmt.Scan(e.ID, e.OrderType, e.SessionID, e.Card, e.EventDate, e.WebSiteURL)
		if err != nil {
			return nil, fmt.Errorf("%s %w", op, err)
		}

		events = append(events, e)
	}

	return events, nil
}
