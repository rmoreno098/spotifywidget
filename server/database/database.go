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
	if err != nil {
		log.Println("Cannot connect to database", sqlite3.ErrAbort)
		return err
	}

	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS users (id TEXT, name TEXT, refresh_token TEXT NOT NULL, access_token TEXT NOT NULL)")
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
	if DB != nil {
		DB.Close()
	}
}

func StoreUserToken(id string, name string, access_token string, refresh_token string) error {
	if DB == nil {
		return errors.New("DATABASE: DB is nil")
	}
	// check if user is already in the db
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM users WHERE id = ?", id).Scan(&count)
	if err != nil {
		log.Println("Database.go Store SELECT COUNT:", err)
		return err
	}

	if count > 0 {
		_, err := DB.Exec("UPDATE users SET access_token = ? WHERE id = ?", access_token, id)
		if err != nil {
			log.Println("Database.go Store UPDATE:", err)
			return err
		}
	} else {
		_, err = DB.Exec("INSERT INTO users (id, name, access_token, refresh_token) VALUES (?, ?, ?, ?)",
		id, name, access_token, refresh_token)
		if err != nil {
			log.Println("Database.go Store INSERT:", err)
			return err
		}
	}
	log.Println("DATABASE: Successfully stored user access_token")
	rows, err := DB.Query("SELECT * FROM users")
	if err != nil {
		log.Println("Database SELECT:", err)
		return err
	}
	for rows.Next() {
		var id, name, access_token, refresh_token string
		if err := rows.Scan(&id, &name, &access_token, &refresh_token); err != nil {
			log.Println("Database Store Printing Select:", err)
		}
		log.Printf("id: %s\n name: %s\n access_token: %s\n refresh_token:%s\n", id, name, access_token, refresh_token)
	}

	return nil
}

func UpdateUserToken(id string, token string) (string, error) {
	if DB == nil {
		return "", errors.New("DATABASE: DB is nil")
	}

	_, err := DB.Exec("UPDATE users SET token = ? WHERE id = ?;", token, id)
	if err != nil {
		return "", err
	}

	return "placeholder", nil
}

func GetUserToken(id string) (string, string, error){
	if DB == nil {
		return "", "", errors.New("DATABASE: DB is nil")
	}

	var access_token, refresh_token string
	err := DB.QueryRow("SELECT access_token, refresh_token FROM users WHERE id = ?", id).Scan(&access_token, &refresh_token)
	if err != nil {
		return "", "", err
	}

	return access_token, refresh_token, nil
}
