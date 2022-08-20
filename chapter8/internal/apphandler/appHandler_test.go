package apphandler

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"

	"log"
	"net/http"
	"testing"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock, err
}

func TestAppHandler(t *testing.T) {
	type args struct {
		m string // method
		r string // route
	}

	tests := map[string]struct {
		args args
		want float64
	}{
		"TestGetEmployee": {args: args{m: http.MethodGet, r: "/employee/1"}, want: http.StatusCreated},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			vars := map[string]string{
				"ID": "1",
			}
			w := httptest.NewRecorder()

			r, _ := http.NewRequest(tt.args.m, tt.args.r, nil)

			r = mux.SetURLVars(r, vars)

			ReadEmployee(w, r)

			assert.Equal(t, tt.want, http.StatusCreated)
		})
	}
}
