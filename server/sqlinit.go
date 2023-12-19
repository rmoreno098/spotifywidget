package main

import (
	"database/sql"
	"log"
	"github.com/mattn/go-sqlite3"
)

func main(){
	log.Println("Initializing SQLite database")
	db, err := sql.Open("sqlite3", "spotifywidget.db")
	if err != nil{
		log.Println("Cannot connect to database", sqlite3.ErrAbort)
		panic(err)
	}
	defer db.Close() // Closes cursor after finished
	err = db.Ping() // Create's a database if does not exist
	if err != nil{
		log.Fatal(err)
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (email VARCHAR(256) PRIMARY KEY, name VARCHAR(64), access_token VARCHAR(255) UNIQUE NOT NULL)")
	if err != nil{
		log.Fatal("Create", err)
		panic(err)
	}

	_, err = db.Exec("INSERT INTO users (email, name, access_token) VALUES (?, ?, ?)", "test@perrohq.com", "Louie Dee", "kjh3%$1212kl3lj5")
	if err != nil{
		log.Fatal("INSERT:", err)
		panic(err)
	}

	rows, err := db.Query("SELECT * FROM users")
	if err != nil{
		log.Fatal("SELECT:", err)
		panic(err)
	}

	for rows.Next(){
		var email, name string
		var token string
		rows.Scan(&email, &name, &token)
		log.Println("Query:", email, name, token)
	}
	log.Println("Query should be 'test@perrohq.com Louie Dee kjh3%$1212kl3lj5'")

}
