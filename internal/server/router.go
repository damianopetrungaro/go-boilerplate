package server

import (
	"net/http"
	"time"

	"github.com/damianopetrungaro/go-boilerplate/pkg/trace"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const traceIDHeader = "X-Trace-ID"

// NewRouter Return a new basic router with some handy middleware
func NewRouter(log logrus.FieldLogger) chi.Router {
	r := chi.NewRouter()

	r.Use(
		profilingMiddleware(log),
		traceMiddleware(),
		middleware.SetHeader("Content-Type", "application/json"),
	)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusMethodNotAllowed)
	})

	return r
}

func profilingMiddleware(log logrus.FieldLogger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rw := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			defer func() {
				start := time.Now()
				log.WithField("status_code", rw.Status()).
					WithField("http_verb", r.Method).
					WithField("bytes", rw.BytesWritten()).
					WithField("latency", time.Since(start).Seconds()).
					WithField("uri", r.URL.String()).
					WithField("trace_id", r.Header.Get(traceIDHeader)).
					Info("router: api call done")
			}()
			next.ServeHTTP(rw, r)
		})
	}
}

func traceMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			traceID := r.Header.Get(traceIDHeader)
			if traceID == "" {
				traceID = uuid.New().String()
			}
			r.WithContext(trace.WithValue(r.Context(), traceID))
			w.Header().Set(traceIDHeader, traceID)
			next.ServeHTTP(w, r)
		})
	}
}
