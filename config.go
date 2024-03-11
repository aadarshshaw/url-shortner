package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	host     = "db"
	port     = 5432
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname   = os.Getenv("DB_NAME")
)

var DB *sql.DB

func OpenConnection() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error
	DB, err = sql.Open("postgres", psqlconn)
	CheckError(err)
	err = DB.Ping()
	CheckError(err)
	qr := `CREATE TABLE IF NOT EXISTS public."Urls" (
		id SERIAL PRIMARY KEY,
		url TEXT NOT NULL,
		shorturl TEXT NOT NULL,
		createdat TIMESTAMP NOT NULL
		);`
	_, err = DB.Exec(qr)
	CheckError(err)

	fmt.Println("Connected!")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
