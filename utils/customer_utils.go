package utils

import (
	"bufio"
	"errors"
	"fmt"
	"golang-database-project/model"
	"os"
	"strconv"
	"strings"
	"time"
)

func ValidateId(id string) error {
	_, err := model.GetCustomerById(id)
	if err != nil {
		return errors.New("customer does't exist")
	}
	return nil
}

func ViewCustomer() {
	customers := model.GetAllCustomer()

	for _, customer := range customers {
		fmt.Printf("%v %v %v %v %s %v \n", customer.Id, customer.Name, customer.Phone, customer.Active_member,
			customer.Join_date.Format("2006-01-02"), customer.Gender)
	}
}

func ViewAllCustomerId() {
	customers := model.GetAllCustomer()

	fmt.Println("AvailableCustomers: ")
	for _, customer := range customers {
		fmt.Printf("%s \n", customer.Id)
	}
}

func DeleteCustomerUtil() {
	var id string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Please enter customer id to be deleted : ")
	scanner.Scan()
	id = scanner.Text()
	err := ValidateId(id)
	if err != nil {
		fmt.Println(err)
	} else {
		model.DeleteCustomer(id)
	}
}

func CreateCustomer() model.Customer {
	var newCustomer model.Customer

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(strings.Repeat("=", 14), "Add New Customer", strings.Repeat("=", 14))
	fmt.Println("Enter Customer Details")
	fmt.Print("Customer Id : ")
	scanner.Scan()
	newCustomer.Id = scanner.Text()

	fmt.Print("Customer Name : ")
	scanner.Scan()
	newCustomer.Name = scanner.Text()

	fmt.Print("Customer Phone Number : ")
	scanner.Scan()
	newCustomer.Phone = scanner.Text()

	fmt.Println()
	fmt.Print("Example: true,t,1,false,f,0\n")
	fmt.Print("Customer Active Member : ")
	scanner.Scan()
	bool_val, _ := strconv.ParseBool(scanner.Text())
	newCustomer.Active_member = bool_val

	fmt.Println()
	fmt.Print("Format yyyy-mm-dd, example: 2023-12-2\n")
	fmt.Print("Customer Join Date : ")
	scanner.Scan()
	fmt.Println(scanner.Text())
	date, _ := time.Parse("2006-01-02", scanner.Text())
	newCustomer.Join_date = date

	fmt.Println()
	fmt.Print("Example: 'F','M','f','m' \n")
	fmt.Print("Customer Gender : ")
	scanner.Scan()
	newCustomer.Gender = strings.ToUpper(scanner.Text())

	return newCustomer
}
