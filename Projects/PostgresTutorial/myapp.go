package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type employee struct {
	fName string
	lName string
}

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "postgres"
	dbname = "postgres"
)

func main()  {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var checkDatabase string
	db.QueryRow("SELECT to_regclass('public.youtube')").Scan(&checkDatabase)
	if err != nil {
		fmt.Print(err)
	}

	if checkDatabase == "" {
		fmt.Println("Creating database")
		createSQL := "CREATE TABLE public.youtube (pk SERIAL PRIMARY KEY,fname character varying,lname character varying);"
		db.Query(createSQL)
	}

	statement :="INSERT INTO youtube(fname, lname) VALUES($1, $2)"

	stmt, err := db.Prepare(statement)
	if err != nil {
		fmt.Print(err)
	}
	defer stmt.Close()

	eName := employee{}

	for i := 0; i < 3; i++ {
		fmt.Print("First Name: ")
		fmt.Scanf("%s", &eName.fName)

		fmt.Print("Last Name: ")
		fmt.Scanf("%s", &eName.lName)
		stmt.QueryRow(eName.fName, eName.lName)
	}

	rows, err := db.Query("SELECT fname, lname from youtube")
	if err != nil {
		fmt.Print(err)
	}
	defer rows.Close()

	fmt.Println("------------------------------------------------")

	for rows.Next() {
		var fname string
		var lname string
		rows.Scan(&fname, &lname)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Printf("%s %s\n", fname, lname)
	}
}