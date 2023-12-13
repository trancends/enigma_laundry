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

func ValidateLength(phone string) error {
	if len(phone) < 11 {
		return errors.New("phone length needs to be at least 11 numbers")
	}

	return nil
}

func ValidateCustomerId(id string) error {
	hasPrefix := strings.HasPrefix(id, "C")
	if !hasPrefix {
		return fmt.Errorf("the first character must be '%s'", "C")
	}
	return nil
}

func DateValidation(date string) (time.Time, error) {
	newDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return time.Time{}, err
	}

	return newDate, nil
}

func CheckCustomerId(id string) error {
	_, err := model.GetCustomerById(id)
	if err != nil {
		return errors.New("customer does't exist")
	}
	return nil
}

func ViewCustomer() {
	customers := model.GetAllCustomer()

	fmt.Println()
	fmt.Println("Customers : ")
	for _, customer := range customers {
		fmt.Printf("%v %v %v %v %s %v \n", customer.Id, customer.Name, customer.Phone, customer.Active_member,
			customer.Join_date.Format("2006-01-02"), customer.Gender)
	}
	fmt.Println()
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
	err := CheckCustomerId(id)
	if err != nil {
		fmt.Println(err)
	} else {
		err := model.DeleteCustomer(id)
		if err != nil {
			fmt.Println("Error", err, "\nThe customer already has relation.")
		}
	}
}

func CreateCustomer() model.Customer {
	var newCustomer model.Customer

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(strings.Repeat("=", 14), "Add New Customer", strings.Repeat("=", 14))
	fmt.Println("Enter Customer Details")

	fmt.Println("Example format for Id 'C001' ")
	for {
		fmt.Print("Customer Id : ")
		scanner.Scan()
		newCustomer.Id = scanner.Text()
		err := ValidateCustomerId(newCustomer.Id)
		if err != nil {
			fmt.Println(err)
		} else {
			break
		}

	}

	fmt.Print("Customer Name : ")
	scanner.Scan()
	newCustomer.Name = scanner.Text()

	for {
		fmt.Print("Customer Phone Number : ")
		scanner.Scan()
		newCustomer.Phone = scanner.Text()
		err := ValidateLength(newCustomer.Phone)
		if err != nil {
			fmt.Println(err)
		} else {
			break
		}
	}

	fmt.Println()
	fmt.Print("Example: true,t,1,false,f,0\n")
	fmt.Print("Customer Active Member : ")
	scanner.Scan()
	bool_val, _ := strconv.ParseBool(scanner.Text())
	newCustomer.Active_member = bool_val

	fmt.Println()
	fmt.Print("Format yyyy-mm-dd, example: 2023-12-2\n")
	for {
		fmt.Print("Customer Join Date : ")
		scanner.Scan()
		fmt.Println(scanner.Text())
		date, err := DateValidation(scanner.Text())
		if err != nil {
			fmt.Println("Please enter a valid Date!")
		} else {
			newCustomer.Join_date = date
			break
		}
	}

	fmt.Println()
	fmt.Print("Example: 'F','M','f','m' \n")
	fmt.Print("Customer Gender : ")
	scanner.Scan()
	newCustomer.Gender = strings.ToUpper(scanner.Text())

	return newCustomer
}
