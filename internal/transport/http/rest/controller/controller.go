package controller

import (
	"net/http"

	"github.com/alexanderiand/notification-service/internal/entity"
)

// Notification UseCase Interface
type UseCase interface {
	NotificationProvider
	NotificationSaver
}

type NotificationProvider interface {
	GetAllEvents() (events []*entity.Event, err error)
}

type NotificationSaver interface {
	SaveEvent(event *entity.Event) (evenID int, err error)
}

// BaseController of the Notification Service
type Controller struct {
	UseCase
}

func New(uc UseCase) *Controller {
	return &Controller{
		UseCase: uc,
	}
}

func (c *Controller) NotifyClient(w http.ResponseWriter, r *http.Request) {
	// TODO: implement

}

// SaveNotification
func (c *Controller) SaveEvent(event *entity.Event) (evenID int, err error) {
	// TODO: implement

	return 0, nil
}

// GetAllEvents()
func (r *Controller) GetAllEvents() (events []*entity.Event, err error) {
	// TODO: implement

	return nil, nil
}
