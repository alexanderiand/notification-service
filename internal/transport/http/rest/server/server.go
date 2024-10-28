package server

import (
	"context"
	"net/http"
	"time"

	"github.com/alexanderiand/notification-service/internal/transport/http/rest/router"
	"github.com/alexanderiand/notification-service/pkg/config"
)

// HTTPServer
type HTTPServer struct {
	httpServer *http.Server
}

// New HTTPServer constructor, return a new instance of the http server
func New(cfg *config.Config, r *router.Router) (srv *HTTPServer) {
	return &HTTPServer{
		httpServer: &http.Server{
			Addr:           cfg.HTTPServer.Addr,
			WriteTimeout:   cfg.HTTPServer.RWTimeout * time.Second, // if 5 * time.Second = 5 second
			ReadTimeout:    cfg.HTTPServer.RWTimeout * time.Second,
			IdleTimeout:    cfg.HTTPServer.RWTimeout * time.Second,
			MaxHeaderBytes: cfg.HTTPServer.MaxHeaderSize << 20, // bitwise ops, if 3 << 20 = equal 3 megabyte
			Handler:        r.Mux,
		},
	}
}

// Start the http server
// If something going wrong, return error
func (s *HTTPServer) Start() error {
	return s.httpServer.ListenAndServe()
}

// Shutdown http server, receive parent context
// If something going wrong, return error
func (s *HTTPServer) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}