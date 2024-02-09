package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/app"
	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/repository"
)

func main() {
	fmt.Println("*** WELCOME to BANKING SYSTEM !! ***")

	//To Initialize Database
	database, err := repository.InitializeDB()
	if err != nil {
		log.Fatalln(err)
	}
	defer database.Close()
	// repository.InsertSeedData(database)

	//Initialize Service
	services := app.NewServices(database)

	//Initialize Router
	router := app.NewRouter(services)

	//Start The Server
	http.ListenAndServe("localhost:1925", router)

}
