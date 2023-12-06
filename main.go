package main

import (
	"fmt"
	"golang-database-project/model"
)

func main() {
	services := model.GetAllService()

	for _, service := range services {
		fmt.Println(service.Id, service.Service_name, service.Satuan, service.Price)
	}
}
