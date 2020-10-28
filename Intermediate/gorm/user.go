package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

type User struct {
	gorm.Model
	Name string
	Email string
}

func initialMigration()  {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err.Error())
		panic("Could not connect to the database")
	}
	defer db.Close()
	db.AutoMigrate(&User{})
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var users []User
	db.Find(&users)

	json.NewEncoder(w).Encode(users)
}

func newUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	db.Create(&User{Name: vars["name"], Email: vars["email"]})
	fmt.Fprintf(w, "New User Successfully Created")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	var user User
	db.Where("name = ?", vars["name"]).Find(&user)
	fmt.Fprintf(w, "Successfully Deleted User")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	var user User
	db.Where("name = ?", vars["name"]).Find(&user)
	user.Email = vars["email"]

	db.Save(&user)
	fmt.Fprintf(w, "Successfully Updated User")
}