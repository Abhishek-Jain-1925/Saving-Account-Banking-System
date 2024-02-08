package main

import (
	"fmt"
	"log"

	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/app"
	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/repository"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("*** WELCOME to BANKING SYSTEM !! ***")

	//To Initialize Database
	database, err := repository.InitializeDB()
	if err != nil {
		log.Fatalln(err)
	}

	//Initialize Service
	services := app.NewServices(database)
	fmt.Println(services)

	//To Routing
	r := mux.NewRouter()
	app.Routes(r, database)

}
