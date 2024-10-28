package middleware

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
)

// Middleware
type BaseMiddleware struct{}

// New is BaseMiddleware constructor, return *BaseMiddleware
func New() *BaseMiddleware {
	return &BaseMiddleware{}
}

// MainMiddleware
func (r *BaseMiddleware) MainMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return r.RequestIDMiddleware(r.LogMiddleware(next))
}

// RequestIDMiddleware for every request generate a new uuid, set for every request this
// uuid into request.Context.Value
// If con't generate a new uuid, return ErrGeneratingRequestID
// If con't get context from request, or con't set the uuid, return ErrRequestContextError
// If something another going wrong, return ErrInternalServerError
func (r *BaseMiddleware) RequestIDMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ruuid := uuid.DefaultGenerator
		uuidV4, err := ruuid.NewV4()
		if err != nil {
			slog.Error(err.Error())
		}

		reqID := uuidV4.String()
		xheader := "X-RequestID"
		w.Header().Set(xheader, reqID)

		log := slog.With(
			slog.String("component", "middleware.RequestIDMiddleware"),
		)
		log.Info("X-RequestID: " + reqID)

		next.ServeHTTP(w, r)
	}
}

// LogMiddleware logged the metadata of the request like remote addr, req time,
func (r *BaseMiddleware) LogMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log := slog.With(
			slog.String("component", "middleware.LogMiddleware"),
		)
		log.Info("logger middleware enabled")

		reqID := w.Header().Get("X-RequestID")

		reqTime := time.Now()
		entry := slog.With(
			slog.String("request_id", reqID),
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.String("remote_addr", r.RemoteAddr),
			slog.String("user_agent", r.UserAgent()),
			slog.String("request_time", reqTime.Format(time.DateTime)),
		)

		next.ServeHTTP(w, r)

		defer func() {
			entry.Info("request completed",
				"request_duration", time.Since(reqTime).String(),
			)
		}()
	}
}

// Other middlewares...
