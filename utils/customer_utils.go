package utils

import (
	"bufio"
	"fmt"
	"golang-database-project/model"
	"os"
	"strconv"
	"strings"
	"time"
)

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
