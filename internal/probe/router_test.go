package probe

import (
	"database/sql"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/go-chi/chi"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {

	tests := []struct {
		scenario string
		fun      func(t *testing.T)
	}{
		{
			scenario: "test routes",
			fun:      testRoutes,
		},
		{
			scenario: "test liveness succed to ping server",
			fun:      testLivenessSucceedToPingServer,
		}, {
			scenario: "test liveness fail to ping server",
			fun:      testLivenessFailToPingServer,
		},
		{
			scenario: "test readiness succed to ping server",
			fun:      testReadinessSucceedToPingServer,
		}, {
			scenario: "test readiness fail to ping server",
			fun:      testReadinessFailToPingServer,
		},
	}

	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			test.fun(t)
		})
	}
}

func testRoutes(t *testing.T) {
	r := chi.NewRouter()
	r.Route("/", NewRouter(&sql.DB{}))
	assert.True(t, r.Match(chi.NewRouteContext(), http.MethodGet, "/liveness"))
	assert.True(t, r.Match(chi.NewRouteContext(), http.MethodGet, "/readiness"))
}

func testLivenessSucceedToPingServer(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.NoError(t, err)
	db.SetMaxIdleConns(1)

	r := liveness(db)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/liveness", nil))

	res := rec.Result()
	assert.Equal(t, http.StatusOK, res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	assert.Equal(t, `["Server is live"]`, string(resBody))

}

func testLivenessFailToPingServer(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.NoError(t, err)
	db.SetMaxIdleConns(0)

	r := liveness(db)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/liveness", nil))

	res := rec.Result()
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	assert.Equal(t, `["Database is not alive"]`, string(resBody))
}

func testReadinessSucceedToPingServer(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.NoError(t, err)
	db.SetMaxIdleConns(1)

	r := readiness(db)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/readiness", nil))

	res := rec.Result()
	assert.Equal(t, http.StatusOK, res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	assert.Equal(t, `["Server is ready"]`, string(resBody))

}

func testReadinessFailToPingServer(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.NoError(t, err)
	db.SetMaxIdleConns(0)

	r := readiness(db)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/readiness", nil))

	res := rec.Result()
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	assert.Equal(t, `["Database is not alive"]`, string(resBody))
}
