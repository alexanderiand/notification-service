package repository

import "github.com/alexanderiand/notification-service/internal/entity"

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

// Base Notification Service Repo
type Repository struct {
	DB Storage
}

// New is Repository constructor, receive Storage interface
// return *Repository
func New(st Storage) *Repository {
	return &Repository{DB: st}
}

// SaveEvent into events table of the sqlite3 database, return  the saved event.id, nil
// If something going wrong, return 0, and error
func (r *Repository) SaveEvent(event *entity.Event) (evenID int, err error) {
	return r.DB.SaveEvent(event)
}

// GetAllEvents select all events from events table of db, return []*entity.Events, nil
// If something going wrong, return nil, and error
func (r *Repository) GetAllEvents() (events []*entity.Event, err error) {
	return r.DB.GetAllEvents()
}
