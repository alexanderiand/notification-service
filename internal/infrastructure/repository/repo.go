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

// SaveNotification
func (r *Repository) SaveEvent(event *entity.Event) (evenID int, err error) {
	// TODO: implement

	return 0, nil
}

// GetAllEvents()
func (r *Repository) GetAllEvents() (events []*entity.Event, err error) {
	// TODO: implement

	return nil, nil
}
