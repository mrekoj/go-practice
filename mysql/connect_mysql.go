package main

import "database/sql"
import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"fmt"
)

func main() {
	db, err := sql.Open("mysql", "root:hd123456@tcp(117.103.207.12:3306)/skynet?parseTime=true")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	// CRUD

	// SELECT
	stmt, err := db.Prepare("SELECT id, name, age, job, created_date FROM users")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var (
		id          int
		name        string
		age         int
		job         string
		createdDate time.Time
	)

	for rows.Next() {
		err := rows.Scan(&id, &name, &age, &job, &createdDate)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("id = %v | name = %v | age = %v | job = %v | time = %v\n", id, name, age, job, createdDate)
		//log.Print(id, name, age, job, createdDate)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
