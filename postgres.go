package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "tharukadevendra"
	password = ""
	dbname   = "golang"
)

func closeDb(db *sql.DB) {
	if err := db.Close(); err != nil {
		panic(err)
	}
	fmt.Println("Successfully closed database")
}

func setupDb() *sql.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}


	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	return db
}
