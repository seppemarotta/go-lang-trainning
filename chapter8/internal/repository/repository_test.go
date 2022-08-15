package repository

import (
	"chapter8/internal/employee"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"log"
	"regexp"
	"testing"
	"time"
)

var e1 = employee.Employee{
	ID:          1,
	FullName:    "Giuseppe Marotta",
	Position:    1,
	Salary:      10,
	Joined:      time.Now(),
	OnProbation: false,
	CreatedAt:   time.Now(),
}

var e2 = employee.Employee{
	ID:          2,
	FullName:    "Giuseppe Clon",
	Position:    2,
	Salary:      20,
	Joined:      time.Now(),
	OnProbation: true,
	CreatedAt:   time.Now(),
}

var e3 = employee.Employee{
	ID:          3,
	FullName:    "Giuseppe Cloncito",
	Position:    2,
	Salary:      500,
	Joined:      time.Now(),
	OnProbation: true,
}

func NewMock() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock, err
}

func TestEmployee_GetId(t *testing.T) {
	db, mock, _ := NewMock()
	repo := &MySqlRepository{db}

	query := regexp.QuoteMeta("SELECT * FROM employees WHERE id = ?")
	prep := mock.ExpectPrepare(query)

	prep.ExpectExec().WithArgs("1").WillReturnResult(sqlmock.NewResult(0, 1))

	user := repo.Get(e1.ID)
	assert.NotNil(t, user)
	//assert.NoError(t, err)
}
func TestEmployee_Insert(t *testing.T) {
	db, mock, _ := NewMock()
	repo := &MySqlRepository{db}

	query := regexp.QuoteMeta("INSERT INTO employees ( FullName, Position, Salary, Joined, OnProbation )\n        VALUES (?, ?, ?, ?, ?)")
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(e1.FullName, e1.Position, e1.Salary, e1.Joined, e1.OnProbation).WillReturnResult(sqlmock.NewResult(0, 1))
	err := repo.Insert(e1.FullName, int(e1.Position), e1.Salary, e1.Joined, e1.OnProbation)
	assert.NoError(t, err)

}
