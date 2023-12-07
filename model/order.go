package model

import (
	"database/sql"
	"fmt"
	"time"
)

type Order struct {
	Id            string
	Date_received time.Time
	Date_finished time.Time
	Customer_id   string
	Receiver      string
}

func AddOrder(order Order) {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := `INSERT INTO order (id, date_received,date_finished, customer_id,receiver) 
    VALUES ($1,$2,$3,$4,$5);`
	result, err := db.Exec(sqlStatement, order.Id, order.Date_received, order.Date_finished, order.Customer_id, order.Receiver)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succesfully Added Order")
	}

	fmt.Println(result)
}

func UpdateOrder(order Order) {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := `UPDATE order SET date_received = $2, date_finished = $3, customer_id = $4 
    receiver = $5 WHERE id = $1`
	result, err := db.Exec(sqlStatement, order.Id, order.Date_received, order.Date_finished, order.Customer_id, order.Receiver)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succesfully Updated Order")
	}

	fmt.Println(result)
}

func GetAllOrder() []Order {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := "SELECT id,Order_name,satuan,price FROM order"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	Orders := scanOrder(rows)
	return Orders
}

func GetOrderById(id string) Order {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := `SELECT * FROM order WHERE id = $1`

	order := Order{}
	err := db.QueryRow(sqlStatement, id).Scan(&order.Id, &order.Date_received, &order.Date_finished,
		&order.Customer_id, &order.Receiver)
	if err != nil {
		panic(err)
	}

	return order
}

func scanOrder(rows *sql.Rows) []Order {
	orders := []Order{}

	for rows.Next() {
		order := Order{}
		err := rows.Scan(&order.Id, &order.Date_received, &order.Date_finished, &order.Customer_id, &order.Receiver)
		if err != nil {
			panic(err)
		}

		orders = append(orders, order)
	}

	err := rows.Err()
	if err != nil {
		panic(err)
	}

	return orders
}
