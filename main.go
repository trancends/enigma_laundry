package main

import (
	"fmt"
	"golang-database-project/model"
)

func main() {
	// customer := entity.Customer{
	// 	Id:            "C005",
	// 	Name:          "Albert Einstein",
	// 	Phone:         "081234789111",
	// 	Active_member: true,
	// 	Join_date:     time.Date(2023, 11, 11, 0, 0, 0, 0, time.Local),
	// 	Gender:        "M",
	// }
	// customer := entity.Customer{
	// 	Id:            "C004",
	// 	Name:          "Albert Einstein",
	// 	Phone:         "081234789111",
	// 	Active_member: true,
	// 	Join_date:     time.Date(2023, 11, 11, 0, 0, 0, 0, time.Local),
	// 	Gender:        "M",
	// }
	// addCustomer(customer)
	// updateCustomer(customer)

	// customers := model.GetAllCustomer()
	// for _, customer := range customers {
	// 	fmt.Println(customer.Id, customer.Name, customer.Phone, customer.Active_member,
	// 		customer.Join_date, customer.Gender)
	// }
	customer := model.GetCustomerById("C001")
	fmt.Println(customer)
}
