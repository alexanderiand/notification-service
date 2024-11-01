package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/alexanderiand/notification-service/internal/transport/http/rest/controller"
	"github.com/alexanderiand/notification-service/internal/transport/http/rest/middleware"
	"github.com/alexanderiand/notification-service/pkg/config"
)

const (
	// endpoint parts
	baseURL = "/api/v1/"
	events  = "events"

	// http methods
	post = "POST "
)

// Router
type Router struct {
	Mux *http.ServeMux
	Ctl *controller.Controller
}

func New(ctl *controller.Controller) *Router {
	return &Router{
		Mux: http.NewServeMux(),
		Ctl: ctl,
	}
}

func (r *Router) InitRouter(cfg *config.Config) {
	// middleware
	mdl := middleware.New()

	// notification-service endpoints

	r.Mux.HandleFunc(endpointJoiner(post, baseURL, events), mdl.MainMiddleware(r.Ctl.NotifyClient))

	// other endpoints...

	// print endpoint in terminal for information user about endpoints
	fmt.Printf("\nNotification Service endpoints:\n\thttp://%s\n\nPress Ctrl+C for stopping service\n",
		endpointJoiner(cfg.HTTPServer.Addr, baseURL, events))
}

// endpoint parts joiner
func endpointJoiner(ep ...string) string {
	return strings.Join(ep, "")
}
