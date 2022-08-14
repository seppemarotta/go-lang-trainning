package apphandler

import (
	"chapter8/internal/employee"
	"chapter8/internal/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router     *mux.Router
	Repository *repository.MySqlRepository
}

var AppInstance App

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
	//e := AppInstance.Repository.Get(id)
	fmt.Fprintf(w, "You've requested the employee: %s.\n", id)
}

func ReadEmployeeByPosition(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]

	fmt.Fprintf(w, "You've requested the employee: %s.\n", id)
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {

	fmt.Println(w, "You've requested the creation of a new employee.")
	var e employee.Employee
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal(reqBody, &e)
	fmt.Println("fullname")

	fmt.Println(e.FullName)
	//AppInstance.Repository.Insert(e.FullName, int(e.Position), e.Salary, e.Joined, e.OnProbation)

	//w.WriteHeader(http.StatusCreated)
	//json.NewEncoder(w).Encode(e)
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]
	fmt.Fprintf(w, "You've requested the employee: %s.\n", id)
}

func Initialize() {
	AppInstance.Repository = repository.InitRepo()
	AppInstance.Router = mux.NewRouter()
	AppInstance.Router.HandleFunc("/", helloWorldHandler)
	AppInstance.Router.HandleFunc("/employee/create", CreateEmployee).Methods("POST")
	AppInstance.Router.HandleFunc("/employee/{ID}", ReadEmployee).Methods("GET")
	AppInstance.Router.HandleFunc("/employee/byposition/{ID}", ReadEmployee).Methods("GET")
	AppInstance.Router.HandleFunc("/employee/{ID}", UpdateEmployee).Methods("PUT")
}

func Run() {
	log.Fatal(http.ListenAndServe(":8080", AppInstance.Router))
}
