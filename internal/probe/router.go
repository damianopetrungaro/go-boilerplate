package probe

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
)

// NewRouter Return a function to use with an existing router
func NewRouter(db *sql.DB) func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/liveness", liveness(db))
		r.Get("/readiness", readiness(db))
	}
}

func liveness(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if err := db.Ping(); err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			_, _ = res.Write([]byte(`["Database is not alive"]`))
			return
		}

		res.WriteHeader(http.StatusOK)
		_, _ = res.Write([]byte(`["Server is live"]`))
	}
}

func readiness(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if err := db.Ping(); err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			_, _ = res.Write([]byte(`["Database is not alive"]`))
			return
		}

		res.WriteHeader(http.StatusOK)
		_, _ = res.Write([]byte(`["Server is ready"]`))
	}
}
