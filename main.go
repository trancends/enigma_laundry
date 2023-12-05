package main

import (
	"database/sql"
	"fmt"
	"golang-database-project/entity"
	"time"

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
	customer := entity.Customer{
		Id:            "C004",
		Name:          "Albert Einstein",
		Phone:         "081234789222",
		Active_member: true,
		Join_date:     time.Date(2023, 11, 11, 0, 0, 0, 0, time.Local),
		Gender:        "M",
	}
	addCustomer(customer)
	// updateCustomer(customer)
	// deleteCustomer("C003")
}

func addCustomer(customer entity.Customer) {
	db := connectDB()
	defer db.Close()

	sqlStatement := `INSERT INTO mst_customer (id, name, phone,active_member,join_date,gender) 
    VALUES ($1,$2,$3,$4,$5,$6);`
	result, err := db.Exec(sqlStatement, customer.Id, customer.Name, customer.Phone,
		customer.Active_member, customer.Join_date, customer.Gender)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succesfully INSERT Data")
	}

	fmt.Println(result)
}

func updateCustomer(customer entity.Customer) {
	db := connectDB()
	defer db.Close()

	sqlStatement := `UPDATE mst_customer SET name = $2, phone = $3, active_member = $4, 
  join_date = $5, gender = $6 WHERE id = $1`
	result, err := db.Exec(sqlStatement, customer.Id, customer.Name, customer.Phone,
		customer.Active_member, customer.Join_date, customer.Gender)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succesfully Updated Data")
	}

	fmt.Println(result)
}

func deleteCustomer(id string) {
	db := connectDB()
	defer db.Close()

	sqlStatement := `DELETE FROM mst_customer WHERE id = $1`
	result, err := db.Exec(sqlStatement, id)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succesfully DELETED Data")
	}

	fmt.Println(result)
}

func connectDB() *sql.DB {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succesfully Connected!")
	}

	return db
}
