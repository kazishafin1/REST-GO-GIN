package usersdb

import (
	"database/sql"
	"log"

	//"github.com/lib/pq" ..
	_ "github.com/lib/pq"
)

var (
	//Client db connection
	Client *sql.DB
	err    error
)

func init() {

	Client, err := sql.Open("postgres", "postgres://postgres:theheart@localhost/mydb?sslmode=disable")

	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("DATABASE CONNECTED")
}
