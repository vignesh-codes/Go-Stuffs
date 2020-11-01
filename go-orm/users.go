package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

var db *gorm.DB

func AllUsers(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could not Connect to Db")

	}
	defer db.Close()

	var users []users
	db.Find(&users)
	json.NewEncoder(w).Encode(users)

	fmt.Fprintf(w, "All Users Endpoint Acessed")
	fmt.Println("All Users Endpoint Acessed")
}

func NewUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New Users Endpoint Acessed")
	fmt.Println("New Users Endpoint Acessed")
}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "delete Users Endpoint Acessed")
	fmt.Println("delete Users Endpoint Acessed")
}

func UpdateUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update Users Endpoint Acessed")
	fmt.Println("All Users Endpoint Acessed")
}
