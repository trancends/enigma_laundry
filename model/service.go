package model

import (
	"database/sql"
	"fmt"
)

type Service struct {
	Id           string
	Service_name string
	Satuan       string
	Price        int
}

func AddService(service Service) {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := `INSERT INTO mst_service (id, service_name, satuan,price) 
    VALUES ($1,$2,$3,$4);`
	result, err := db.Exec(sqlStatement, service.Id, service.Service_name, service.Satuan, service.Price)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succesfully INSERT Service")
	}

	fmt.Println(result)
}

func UpdateService(service Service) {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := `UPDATE mst_service SET service_name = $2, satuan = $3, price = $4 WHERE id = $1`
	result, err := db.Exec(sqlStatement, service.Id, service.Service_name, service.Satuan, service.Price)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succesfully Update Service")
	}

	fmt.Println(result)
}

func DeleteService(id string) {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := `DELETE FROM mst_service WHERE id = $1`
	result, err := db.Exec(sqlStatement, id)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succesfully DELETE Service")
	}

	fmt.Println(result)
}

func GetAllService() []Service {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := "SELECT id,service_name,satuan,price FROM mst_service"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	services := scanService(rows)
	return services
}

func scanService(rows *sql.Rows) []Service {
	services := []Service{}

	for rows.Next() {
		service := Service{}
		err := rows.Scan(&service.Id, &service.Service_name, &service.Satuan, &service.Price)
		if err != nil {
			panic(err)
		}

		services = append(services, service)
	}

	err := rows.Err()
	if err != nil {
		panic(err)
	}

	return services
}
