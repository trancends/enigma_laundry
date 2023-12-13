package model

import (
	"database/sql"
	"fmt"
)

type OrderDetail struct {
	Id         string
	Order_id   string
	Service_id string
	Quantity   int
}

func AddOrderDetail(orderDetail OrderDetail) error {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := `INSERT INTO order_detail (id,order_id,service_id,quantity) 
    VALUES ($1,$2,$3,$4);`
	result, err := db.Exec(sqlStatement, orderDetail.Id, orderDetail.Order_id,
		orderDetail.Service_id, orderDetail.Quantity)

	if err != nil {
		return err
	} else {
		fmt.Println("Succesfully Added OrderDetail")
	}

	fmt.Println(result)
	return err
}

func UpdateOrderDetail(orderDetail OrderDetail) error {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := `UPDATE order_detail SET order_id = $2, service_id = $3, quantity = $4 WHERE id = $1`
	result, err := db.Exec(sqlStatement, orderDetail.Id, orderDetail.Order_id,
		orderDetail.Service_id, orderDetail.Quantity)

	if err != nil {
		return err
	} else {
		fmt.Println("Succesfully Updated OrderDetail")
	}

	fmt.Println(result)
	return err
}

func GetAllOrderDetail() []OrderDetail {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := "SELECT id,order_id,service_id,quantity FROM order_detail"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	orderDetails := scanOrderDetail(rows)
	return orderDetails
}

func GetOrderDetailById(id string) (OrderDetail, error) {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := `SELECT * FROM order_detail WHERE id = $1`

	orderDetail := OrderDetail{}
	err := db.QueryRow(sqlStatement, id).Scan(&orderDetail.Id, &orderDetail.Order_id,
		&orderDetail.Service_id, &orderDetail.Quantity)

	return orderDetail, err
}

func scanOrderDetail(rows *sql.Rows) []OrderDetail {
	orderDetails := []OrderDetail{}

	for rows.Next() {
		orderDetail := OrderDetail{}
		err := rows.Scan(&orderDetail.Id, &orderDetail.Order_id,
			&orderDetail.Service_id, &orderDetail.Quantity)
		if err != nil {
			panic(err)
		}

		orderDetails = append(orderDetails, orderDetail)
	}

	err := rows.Err()
	if err != nil {
		panic(err)
	}

	return orderDetails
}
