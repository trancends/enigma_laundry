package model

import (
	"database/sql"
	"fmt"
	"time"
)

type Customer struct {
	Id            string
	Name          string
	Phone         string
	Active_member bool
	Join_date     time.Time
	Gender        string
}

func AddCustomer(customer Customer) error {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := `INSERT INTO mst_customer (id, name, phone,active_member,join_date,gender) 
    VALUES ($1,$2,$3,$4,$5,$6);`
	result, err := db.Exec(sqlStatement, customer.Id, customer.Name, customer.Phone,
		customer.Active_member, customer.Join_date, customer.Gender)

	if err != nil {
		return err
	} else {
		fmt.Println("Succesfully Added Customer")
	}

	fmt.Println(result)
	return err
}

func UpdateCustomer(customer Customer) error {
	db := ConnectDB()
	defer db.Close()

	_, err := GetCustomerById(customer.Id)
	if err != nil {
		return err
	}

	sqlStatement := `UPDATE mst_customer SET name = $2, phone = $3, active_member = $4, 
  join_date = $5, gender = $6 WHERE id = $1`
	result, err := db.Exec(sqlStatement, customer.Id, customer.Name, customer.Phone,
		customer.Active_member, customer.Join_date, customer.Gender)

	if err != nil {
		return err
	} else {
		fmt.Println("Succesfully Updated Customer")
	}

	fmt.Println(result)
	return nil
}

func DeleteCustomer(id string) error {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := `DELETE FROM mst_customer WHERE id = $1`
	result, err := db.Exec(sqlStatement, id)

	if err != nil {
		return err
	} else {
		fmt.Println("Succesfully DELETED Customer")
	}

	fmt.Println(result)
	return err
}

func GetAllCustomer() []Customer {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := "SELECT * FROM mst_customer"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	customers := scanCustomer(rows)
	return customers
}

func GetCustomerById(id string) (Customer, error) {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := `SELECT * FROM mst_customer WHERE id = $1`

	customer := Customer{}
	err := db.QueryRow(sqlStatement, id).Scan(&customer.Id, &customer.Name, &customer.Phone,
		&customer.Active_member, &customer.Join_date, &customer.Gender)

	return customer, err
}

func scanCustomer(rows *sql.Rows) []Customer {
	customers := []Customer{}

	for rows.Next() {
		customer := Customer{}
		err := rows.Scan(&customer.Id, &customer.Name, &customer.Phone, &customer.Active_member,
			&customer.Join_date, &customer.Gender)
		if err != nil {
			panic(err)
		}

		customers = append(customers, customer)
	}

	err := rows.Err()
	if err != nil {
		panic(err)
	}

	return customers
}

func searchBy(name string) []Customer {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := `SELECT * FROM mst_customer WHERE name LIKE $1`

	rows, err := db.Query(sqlStatement, "%"+name+"%")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	customers := scanCustomer(rows)

	return customers
}
