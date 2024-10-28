package usecase

import "github.com/alexanderiand/notification-service/internal/entity"

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
	// TODO: implement

	return nil, nil
}

func (uc *UseCase) SaveEvent(event *entity.Event) (eventID int, err error) {
	// TODO: implement

	return 0, nil
}
