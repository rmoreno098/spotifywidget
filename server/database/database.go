package database

import (
	"database/sql"
	"errors"
	"log"

	"github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() error {
	log.Println("Initializing SQLite database")
	var err error
	DB, err = sql.Open("sqlite3", "./spotifywidget.db")
	if err != nil{
		log.Println("Cannot connect to database", sqlite3.ErrAbort)
		return err
	}

	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS users (id TEXT, name TEXT, token TEXT)")
	if err != nil {
		log.Println("Cannot create table", sqlite3.ErrAbort)
		return err
	}

	err = DB.Ping()
    if err != nil {
        log.Fatal(err)
		return err
    }

	return nil
}

func CloseDB() {
	if DB != nil{
		DB.Close()
	}
}

func StoreUsrToken(id string, name string, token string) error {
	if DB == nil {
		log.Println("DATABASE: DB is nil")
		return errors.New("DATABASE: DB is nil")
	}

	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM users WHERE id = ?", id).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("DATABASE: User already exists")
	}

	_, err = DB.Exec("INSERT INTO users (id, name, token) VALUES (?, ?, ?)", id, name, token)
	if err != nil {
		return err
	}

	log.Println("DATABASE: Successfully stored user token")
	rows, err := DB.Query("SELECT * FROM users")
	if err != nil {
		return err
	}
	for rows.Next() {
		var id, name, token string
		if err := rows.Scan(&id, &name, &token); err != nil {
			log.Println(err)
			continue
		}
		log.Println("DATABASE:", id, name, token)
	}
	defer rows.Close()
	
	return nil
}