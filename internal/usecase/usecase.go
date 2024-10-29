package usecase

import (
	"log/slog"

	"github.com/alexanderiand/notification-service/internal/entity"
)

// main Repository interface
type NotificationRepo interface {
	NotificationProvider
	NotificationSaver
}

type NotificationProvider interface {
	GetAllEvents() (events []*entity.Event, err error)
}

type NotificationSaver interface {
	SaveEvent(event *entity.Event) (evenID int, err error)
}

// Base Notification Service usecase
type UseCase struct {
	NotificationRepo
}

// New is NotificationUseCase constructor, receive interface, return *UseCase
func New(nuc NotificationRepo) *UseCase {
	return &UseCase{NotificationRepo: nuc}
}

func (uc *UseCase) GetAllEvents() (events []*entity.Event, err error) {
	// other app logic

	events, err = uc.NotificationRepo.GetAllEvents()
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	return events, nil
}

// SaveEvent into events table of the sqlite3 database, return  the saved event.id, nil
// If something going wrong, return 0, and error
func (uc *UseCase) SaveEvent(event *entity.Event) (eventID int, err error) {
	// other app logic

	eventID, err = uc.NotificationRepo.SaveEvent(event)
	if err != nil {
		slog.Error(err.Error())
		return 0, err
	}

	return eventID, nil
}
