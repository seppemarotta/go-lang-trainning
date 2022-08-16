package apphandler

import (
	"chapter8/internal/employee"
	"chapter8/internal/repository"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type App struct {
	Router     *mux.Router
	Repository *repository.MySqlRepository
}

var AppInstance *App

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	var response map[string]interface{}
	json.Unmarshal([]byte(`{ "hello": "world" }`), &response)
	respondWithJSON(w, http.StatusOK, response)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

func ReadEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]
	parsedId, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	e, err := AppInstance.Repository.Employee(context.Background(), parsedId)
	fmt.Fprintf(w, "You've requested the employee: %s.\n", id)
	json.NewEncoder(w).Encode(e)
}

func ReadEmployeeByPos(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]
	parsedId, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	e, err := AppInstance.Repository.EmployeeByPos(parsedId)
	fmt.Fprintf(w, "You've requested the employee: %s.\n", id)
	json.NewEncoder(w).Encode(e)
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "You've requested the creation of a new employee.")
	var e employee.Employee

	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	AppInstance.Repository.Save(context.Background(), &e)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(e)
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]
	parsedId, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	fmt.Println(w, "You've requested the update of a new employee.")
	var e employee.Employee

	err = json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	AppInstance.Repository.Update(parsedId, e.FullName, int(e.Position), e.Salary, e.Joined, e.OnProbation)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(e)
	fmt.Fprintf(w, "You've requested the employee: %s.\n", id)
}

func Initialize(rep *repository.MySqlRepository) (a *App) {
	AppInstance.Repository = rep
	AppInstance.Router = mux.NewRouter()
	AppInstance.Router.HandleFunc("/", helloWorldHandler)
	AppInstance.Router.HandleFunc("/employee/create", CreateEmployee).Methods("POST")
	AppInstance.Router.HandleFunc("/employee/{ID}", ReadEmployee).Methods("GET")
	AppInstance.Router.HandleFunc("/employee/byposition/{ID}", ReadEmployeeByPos).Methods("GET")
	AppInstance.Router.HandleFunc("/employee/{ID}", UpdateEmployee).Methods("PUT")
	return a
}

func Run() {
	log.Fatal(http.ListenAndServe(":8080", AppInstance.Router))
}
