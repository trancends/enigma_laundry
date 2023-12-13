package utils

import (
	"bufio"
	"errors"
	"fmt"
	"golang-database-project/model"
	"os"
	"strings"
)

func ValidateOrderId(id string) error {
	hasPrefix := strings.HasPrefix(id, "S")
	if !hasPrefix {
		return fmt.Errorf("the first character must be '%s'", "O")
	}
	return nil
}

func CheckOrderId(id string) error {
	_, err := model.GetOrderById(id)
	if err != nil {
		return errors.New("customer does't exist")
	}
	return nil
}

func ViewOrder() {
	orders := model.GetAllOrder()

	fmt.Println()
	fmt.Println("Services : ")
	for _, order := range orders {
		fmt.Printf("%s, %s, %s, %s, %s\n", order.Id, order.Date_received, order.Date_finished,
			order.Customer_id, order.Receiver)
	}
	fmt.Println()
}

func ViewAllorderId() {
	orders := model.GetAllOrder()

	fmt.Println("Availableorders: ")
	for _, order := range orders {
		fmt.Printf("%s \n", order.Id)
	}
}

func CreateOrder() model.Order {
	var newOrder model.Order

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(strings.Repeat("=", 14), "Add New order", strings.Repeat("=", 14))
	fmt.Println("Enter order Details")

	fmt.Println("Example format for Id 'S001'")
	for {
		fmt.Print("order Id : ")
		scanner.Scan()
		newOrder.Id = scanner.Text()
		err := ValidateOrderId(newOrder.Id)
		if err != nil {
			fmt.Println(err)
		} else {
			break
		}

	}

	fmt.Println()
	fmt.Print("Format yyyy-mm-dd, example: 2023-12-2\n")
	for {
		fmt.Print("Order Received Date : ")
		scanner.Scan()
		fmt.Println(scanner.Text())
		date, err := DateValidation(scanner.Text())
		if err != nil {
			fmt.Println("Please enter a valid Date!")
		} else {
			newOrder.Date_received = date
			break
		}
	}

	for {
		fmt.Print("Order Finished Date : ")
		scanner.Scan()
		fmt.Println(scanner.Text())
		date, err := DateValidation(scanner.Text())
		if err != nil {
			fmt.Println("Please enter a valid Date!")
		} else {
			newOrder.Date_finished = date
			break
		}
	}

	fmt.Println("Example format for Id 'C001' ")
	for {
		fmt.Print("Customer Id : ")
		scanner.Scan()
		newOrder.Customer_id = scanner.Text()
		err := ValidateId(newOrder.Customer_id)
		if err != nil {
			fmt.Println(err)
		} else {
			break
		}

	}

	fmt.Print("Receiver : ")
	scanner.Scan()
	newOrder.Receiver = scanner.Text()

	return newOrder
}
