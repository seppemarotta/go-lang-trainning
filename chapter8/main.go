package main

import (
	"chapter8/internal/apphandler"
	"fmt"
)

/*
	func handleGet(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "get\n")
		//by id & pos
	}

	func handlePost(w http.ResponseWriter, req *http.Request) {
		//create
		rep := repository.InitRepo()
		e := employee.Employee{}
		fmt.Fprintf(w, "post\n")
		reqBody, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
		}

		json.Unmarshal(reqBody, &e)

		rep.Insert(e.FullName, int(e.Position), e.Salary, e.Joined, e.OnProbation)

		fmt.Println(e)
	}

	func handlePut(w http.ResponseWriter, req *http.Request) {
		//create
		fmt.Fprintf(w, "post\n")
	}

	func reqHandler(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		fmt.Printf("got / request\n")
		//io.WriteString(w, "This is my website!\n")
	}
*/
func main() {
	fmt.Println("Good")
	//rep := repository.InitRepo()
	//e:=rep.Get(1)
	//fmt.Println(e)
	//rep.Insert("Giuseppe Marotta", 1, 123.123, time.Now(), true)
	//rep.Update(3,"Giuseppe Test",0,1.1,time.Now(),true)
	//rep.EmployeeByPos(1)

	apphandler.Initialize()
	apphandler.Run()

}
