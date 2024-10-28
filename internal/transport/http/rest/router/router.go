package router

import (
	"net/http"

	"github.com/alexanderiand/notification-service/internal/transport/http/rest/controller"
)

// Router
type Router struct {
	Mux *http.ServeMux
	Ctl *controller.Controller
	// Middleware
}

func New(ctl *controller.Controller) *Router {
	return &Router{
		Mux: http.NewServeMux(),
		Ctl: ctl,
	}
}

func (r *Router) InitRouter() {
	// middleware

	// notification-service endpoints

	r.Mux.HandleFunc("POST /events", r.Ctl.NotifyClient)

	// other endpoints...

}
