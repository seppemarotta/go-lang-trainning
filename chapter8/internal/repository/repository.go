package repository

import (
	"chapter8/internal/db"
	e "chapter8/internal/employee"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySqlRepository struct {
	db *sql.DB
}

func InitRepo() *MySqlRepository {
	return &MySqlRepository{db: db.InitDB()}
}

func (repository *MySqlRepository) Get(ID int) e.Employee {
	query, err := repository.db.Prepare("Select * from employees where id = ?")
	if err != nil {
		log.Fatal(err.Error())
	}
	employee := e.Employee{}
	err = query.QueryRow(ID).Scan(&employee.ID, &employee.FullName, &employee.Position, &employee.Salary, &employee.Joined, &employee.OnProbation, &employee.CreatedAt)
	if err != nil {
		log.Fatal(err.Error())
	}
	return employee
}

func (repository *MySqlRepository) Insert(fullName string, pos int, salary float64, joined time.Time, probation bool) {
	sqlStatement := `
	INSERT INTO employees ( FullName, Position, Salary, Joined, OnProbation )
        VALUES (?, ?, ?, ?, ?)`
	_, err := repository.db.Exec(sqlStatement, fullName, pos, salary, joined, probation)
	if err != nil { // scan will release the connection
		log.Fatal("Error")
		log.Fatal(err.Error())
		panic(err)
	}
}

func (repository *MySqlRepository) Update(ID int, fullName string, pos int, salary float64, joined time.Time, probation bool) {
	sqlStatement := `
	UPDATE employees set FullName = ?, Position = ?, Salary = ?, Joined = ?, OnProbation = ? where ID = ?
    `
	_, err := repository.db.Exec(sqlStatement, ID, fullName, pos, salary, joined, probation)
	if err != nil { // scan will release the connection
		log.Fatal("Error")
		log.Fatal(err.Error())
		panic(err)
	}
}

func (repository *MySqlRepository) EmployeeByPos(pos int) {
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
			//e:= Employee{}

		}
	}
}
