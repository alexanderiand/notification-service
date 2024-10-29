package controller

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/alexanderiand/notification-service/internal/entity"
)

// worker pool
const (
	workerCount  = 3
	jobQueueSize = 10
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
	ev := entity.Event{}

	if err := json.NewDecoder(r.Body).Decode(&ev); err != nil {
		slog.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	ev.CreatedAt = time.Now()

	id, err := c.UseCase.SaveEvent(&ev)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	slog.Debug("event successful saved in the database with", "eventID", strconv.Itoa(id))
	ev.ID = id

	// response with http status code 201 - StatusCreated
	w.WriteHeader(http.StatusCreated)

	// job channel for workers
	jobChan := make(chan entity.Event, jobQueueSize)

	select {
	case jobChan <- ev:
		slog.Debug("job send to the jobChan")
	default:
		slog.Warn("the jobChan was overflow")
		http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
	}

	go c.WorkerPool(workerCount, jobChan)
}

// SaveEvent into events table of the sqlite3 database, return  the saved event.id, nil
// If something going wrong, return 0, and error
func (c *Controller) SaveEvent(event *entity.Event) (evenID int, err error) {
	// other controller logic
	return c.UseCase.SaveEvent(event)
}

// GetAllEvents select all events from events table of db, return []*entity.Events, nil
// If something going wrong, return nil, and error
func (r *Controller) GetAllEvents() (events []*entity.Event, err error) {
	// other controller logic

	return r.UseCase.GetAllEvents()
}

// WorkerPool simple worker pool for sending notification to client for event
func (c *Controller) WorkerPool(workerCount int, jobQueue <-chan entity.Event) {
	wg := sync.WaitGroup{}

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for event := range jobQueue {
				fmt.Printf("\nSending notification for event: \n\t%+v\n\n", event) //
				time.Sleep(time.Second)                                            // sending notification...
			}
		}(i)
	}
	wg.Wait()
}
