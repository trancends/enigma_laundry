package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "enigmacamp"
	password = "1234"
	dbname   = "enigma_laundry"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func main() {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succesfully Connected!")
	}

	sqlStatement := `INSERT INTO mst_customer (id, name, phone,active_member,join_date)
    VALUES ('C001','Sari Ayunda','081234567999','true','2020-02-23')`
	result, err := db.Exec(sqlStatement)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succesfully INSERT Data")
	}

	fmt.Println(result)
}

func addCustomer(id, name, gender, phone, active_member, join_date string)
