package repository

import (
	e "chapter8/internal/employee"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type Repository interface {
	Employees(ctx context.Context, pos e.Position) ([]e.Employee, error)
	Employee(ctx context.Context, id int) (*e.Employee, error)
	Save(ctx context.Context, e *e.Employee) error
}

type MySqlRepository struct {
	db *sql.DB
}

func InitRepo(db *sql.DB) *MySqlRepository {
	return &MySqlRepository{db: db}
}

func (repository *MySqlRepository) Employee(ctx context.Context, id int) (*e.Employee, error) {
	query, err := repository.db.Prepare("SELECT * FROM employees WHERE id = ?")
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	employee := e.Employee{}
	err = query.QueryRow(id).Scan(&employee.ID, &employee.FullName, &employee.Position, &employee.Salary, &employee.Joined, &employee.OnProbation, &employee.CreatedAt)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return &employee, nil
}

func (repository *MySqlRepository) Save(ctx context.Context, e *e.Employee) error {
	sqlStatement := `INSERT INTO employees ( FullName, Position, Salary, Joined, OnProbation ) VALUES (?, ?, ?, ?, ?) RETURNING employees.ID`
	_, err := repository.db.Exec(sqlStatement, e.FullName, e.Position, e.Salary, e.Joined, e.OnProbation)
	if err != nil { // scan will release the connection
		log.Fatal(err.Error())
		return err
	}
	return nil
}

func (repository *MySqlRepository) Update(ID int, fullName string, pos int, salary float64, joined time.Time, probation bool) {
	sqlStatement := `
	UPDATE employees set FullName = ?, Position = ?, Salary = ?, Joined = ?, OnProbation = ? where ID = ?
    `
	_, err := repository.db.Exec(sqlStatement, fullName, pos, salary, joined, probation, ID)
	if err != nil { // scan will release the connection
		log.Fatal("Error")
		log.Fatal(err.Error())
		panic(err)
	}
}

func (repository *MySqlRepository) EmployeeByPos(pos int) (employees []e.Employee, err error) {
	sqlStatement := `
	SELECT * FROM employees where Position = ?
    `

	rows, err := repository.db.Query(sqlStatement, pos)
	if err != nil { // scan will release the connection
		log.Fatal("Error")
		log.Fatal(err.Error())
		panic(err)
	} else {
		defer rows.Close()
		for rows.Next() {
			e := e.Employee{}

			rows.Scan(&e.ID, &e.FullName, &e.Position, &e.Salary, &e.Joined, &e.OnProbation, &e.CreatedAt)
			fmt.Println(e)
			employees = append(employees, e)
		}
	}
	return employees, nil
}
