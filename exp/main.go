package main

import (
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "lenslocked_dev"
	password = "lenslocked_dev"
	dbname = "lenslocked_dev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	defer db.Close()

	var id int

	err = db.QueryRow(`
	INSERT INTO users(name,email)
	VALUES ($1, $2) RETURNING id`,"Stoyan Stoyanov","melloboy89@gmail.com").Scan(&id)

	if err!=nil{
		panic(err)
	}

	fmt.Println("id is ...", id)
}
