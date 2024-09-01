package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func InitDb() {
	dbEnv, err := getDbEnv()

	if err != nil {
		panic(err)
	}

	db, err := sql.Open("sqlite3", dbEnv.DB_PATH)

	if err != nil {
		panic("Could not connect to database")
	}

	db.SetMaxOpenConns(dbEnv.DB_MAX_OPEN_CONNS)
	db.SetMaxIdleConns(dbEnv.DB_MAX_IDLE_CONNS)
	Db = db

	err = createTables()

	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to the db")
}

func createTables() error {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)`

	_, err := Db.Exec(createEventsTable)

	if err != nil {
		return fmt.Errorf(`error: Couldn't create events table, err: %v`, err)
	}

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)`

	_, err = Db.Exec(createUsersTable)

	if err != nil {
		return fmt.Errorf(`error: Couldn't create users table, err: %v`, err)
	}
	return nil
}

type DbEnv struct {
	DB_PATH           string
	DB_MAX_OPEN_CONNS int
	DB_MAX_IDLE_CONNS int
}

func getDbEnv() (DbEnv, error) {
	DB_PATH, varPresent := os.LookupEnv("DB_PATH")

	if !varPresent {
		return DbEnv{}, errors.New("error: DB_PATH .env variable not found")
	}

	DB_MAX_OPEN_CONNS_STR, varPresent := os.LookupEnv("DB_MAX_OPEN_CONNS")

	if !varPresent {
		return DbEnv{}, errors.New("error: DB_MAX_OPEN_CONNS .env variable not found")
	}

	DB_MAX_IDLE_CONNS_STR, varPresent := os.LookupEnv("DB_MAX_IDLE_CONNS")

	if !varPresent {
		return DbEnv{}, errors.New("error: DB_MAX_IDLE_CONNS .env variable not found")
	}

	DB_MAX_OPEN_CONNS, parseIntErr := strconv.Atoi(DB_MAX_OPEN_CONNS_STR)

	if parseIntErr != nil {
		return DbEnv{}, errors.New("error: DB_MAX_OPEN_CONNS .env variable couldn't be parsed into int")
	}

	DB_MAX_IDLE_CONNS, parseIntErr := strconv.Atoi(DB_MAX_IDLE_CONNS_STR)

	if parseIntErr != nil {
		return DbEnv{}, errors.New("error: DB_MAX_IDLE_CONNS_STR .env variable couldn't be parsed into int")
	}

	dbEnv := DbEnv{DB_PATH, DB_MAX_OPEN_CONNS, DB_MAX_IDLE_CONNS}

	return dbEnv, nil
}
