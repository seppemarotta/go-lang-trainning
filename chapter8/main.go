package main

import (
	"chapter8/internal/apphandler"
	"chapter8/internal/db"
	"chapter8/internal/repository"
	"fmt"
)

func main() {
	fmt.Println("Good")
	r := repository.InitRepo(db.InitDB())
	apphandler.Initialize(r)
	apphandler.Run()
}
