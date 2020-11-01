package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	fmt.Println("MySQL basic")

	db, err := sql.Open("mysql", "root:vignesh123@tcp(127.0.0.1:3306)/testdb")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	fmt.Println("Successfully Connected to DB")

	// insert, err := db.Query("INSERT INTO users VALUES('VicSSky')")

	// if err != nil {
	// 	panic(err.Error())
	// }

	// defer insert.Close()
	// fmt.Println("Successfully inserted into DB")
	results, err := db.Query("SELECT name FROM users")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var user User

		err = results.Scan(&user.Name)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println(user.Name)

	}

	// defer results.Close()
	// fmt.Println("Successfully resultsed into DB")

}
