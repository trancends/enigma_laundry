package model

import (
	"database/sql"
	"fmt"
	"golang-database-project/entity"
)

func AddCustomer(customer entity.Customer) {
	db := ConnectDB()
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

func UpdateCustomer(customer entity.Customer) {
	db := ConnectDB()
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

func DeleteCustomer(id string) {
	db := ConnectDB()
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

func GetAllCustomer() []entity.Customer {
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

func GetCustomerById(id string) entity.Customer {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := `SELECT * FROM mst_customer WHERE id = $1`

	customer := entity.Customer{}
	err := db.QueryRow(sqlStatement, id).Scan(&customer.Id, &customer.Name, &customer.Phone,
		&customer.Active_member, &customer.Join_date, &customer.Gender)
	if err != nil {
		panic(err)
	}

	return customer
}

func scanCustomer(rows *sql.Rows) []entity.Customer {
	customers := []entity.Customer{}

	for rows.Next() {
		customer := entity.Customer{}
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

func searchBy(name string) []entity.Customer {
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
