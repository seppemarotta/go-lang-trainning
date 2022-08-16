package repository

import (
	"chapter8/internal/employee"
	"context"
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
	ctx := context.Background()
	query := regexp.QuoteMeta("SELECT * FROM employees WHERE id = ?")

	rows := sqlmock.NewRows([]string{"id", "full_name", "position", "salary", "joined", "on_probation", "created_at"}).
		AddRow(e1.ID, e1.FullName, e1.Position, e1.Salary, e1.Joined, e1.OnProbation, e1.CreatedAt)

	mock.ExpectPrepare(query).ExpectQuery().WithArgs(1).WillReturnRows(rows)

	user, err := repo.Employee(ctx, e1.ID)
	assert.NotNil(t, user)
	assert.NoError(t, err)

}

func TestEmployee_Insert(t *testing.T) {
	db, mock, _ := NewMock()
	ctx := context.Background()
	repo := &MySqlRepository{db}

	//query := regexp.QuoteMeta("INSERT INTO employees ( FullName, Position, Salary, Joined, OnProbation ) VALUES (?, ?, ?, ?, ?)")
	query := regexp.QuoteMeta(`INSERT INTO employees ( FullName, Position, Salary, Joined, OnProbation ) VALUES (?, ?, ?, ?, ?) RETURNING employees.ID`)
	//mock.ExpectBegin()
	mock.ExpectPrepare(query).ExpectQuery().WithArgs(e1.FullName, e1.Position, e1.Salary, e1.Joined, e1.OnProbation).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	err := repo.Save(ctx, &e1)

	assert.NoError(t, err)
}
