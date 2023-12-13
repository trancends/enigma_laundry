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

func AddOrder(order Order) error {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := `INSERT INTO orders (id, date_received,date_finished, customer_id,receiver) 
    VALUES ($1,$2,$3,$4,$5);`
	result, err := db.Exec(sqlStatement, order.Id, order.Date_received, order.Date_finished, order.Customer_id, order.Receiver)

	if err != nil {
		return err
	} else {
		fmt.Println("Succesfully Added Order")
	}

	fmt.Println(result)
	return err
}

func UpdateOrder(order Order) error {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := `UPDATE orders SET date_received = $2, date_finished = $3, customer_id = $4,
    receiver = $5 WHERE id = $1`
	result, err := db.Exec(sqlStatement, order.Id, order.Date_received, order.Date_finished, order.Customer_id, order.Receiver)

	if err != nil {
		return err
	} else {
		fmt.Println("Succesfully Updated Order")
	}

	fmt.Println(result)
	return err
}

func GetAllOrder() []Order {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := "SELECT * FROM orders"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	Orders := scanOrder(rows)
	return Orders
}

func GetOrderById(id string) (Order, error) {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := `SELECT * FROM orders WHERE id = $1`

	order := Order{}
	err := db.QueryRow(sqlStatement, id).Scan(&order.Id, &order.Date_received, &order.Date_finished,
		&order.Customer_id, &order.Receiver)

	return order, err
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
