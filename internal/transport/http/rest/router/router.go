package router

import "net/http"

// Router
type Router struct {
	Mux *http.ServeMux
	// Controller
	// Middleware
}

func New() *Router {
	return &Router{
		// Ctl, Mdl
	}
}

func (r *Router) InitRouter() {
	// middleware

	// notification-service endpoints

	r.Mux.HandleFunc("POST /events", r.Ctl.EventNotifier)

	// other endpoints...

}