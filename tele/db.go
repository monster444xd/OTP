package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go MySQL Tutorial bvm7x")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "root:shl@tcp(127.0.0.1:3306)/bypassbot")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO `bypassbot`.`users` (`idTelagramID`, `DPS`, `TelagramID`, `Administrator`, `BlockedUser`, `expiry`) VALUES ('3', '3', '3', '3', '3', '3');")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()

}
