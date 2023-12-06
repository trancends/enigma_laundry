package main

import (
	"fmt"
	"golang-database-project/model"
)

func main() {
	service := model.Service{
		Id:           "S004",
		Service_name: "Laundry Boneka",
		Satuan:       "Buah",
		Price:        25000,
	}
	fmt.Println(service)
	model.AddService(service)
}
