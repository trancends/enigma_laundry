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

func AddService(service Service) error {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := `INSERT INTO mst_service (id, service_name, satuan,price) 
    VALUES ($1,$2,$3,$4);`
	result, err := db.Exec(sqlStatement, service.Id, service.Service_name, service.Satuan, service.Price)

	if err != nil {
		return err
	} else {
		fmt.Println("Succesfully Added Customer")
	}

	fmt.Println(result)
	return err
}

func UpdateService(service Service) error {
	db := ConnectDB()
	defer db.Close()

	_, err := GetServiceById(service.Id)
	if err != nil {
		return err
	}

	sqlStatement := `UPDATE mst_service SET service_name = $2, satuan = $3, price = $4 WHERE id = $1`
	result, err := db.Exec(sqlStatement, service.Id, service.Service_name, service.Satuan, service.Price)

	if err != nil {
		return err
	} else {
		fmt.Println("Succesfully Updated Service")
	}

	fmt.Println(result)
	return nil
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

func GetServiceById(id string) (Service, error) {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := `SELECT * FROM mst_service WHERE id = $1`

	service := Service{}
	err := db.QueryRow(sqlStatement, id).Scan(&service.Id, &service.Service_name,
		&service.Satuan, &service.Price)

	return service, err
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
