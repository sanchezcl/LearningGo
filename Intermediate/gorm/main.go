package main

import (
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"net/http"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", allUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", updateUser).Methods("PUT")
	myRouter.HandleFunc("/user/{name}/{email}", newUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	log.Println("GO ORM tutorial")
	log.Println("Server start!")
	initialMigration()
	handleRequests()
}
