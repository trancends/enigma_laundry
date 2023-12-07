package main

import (
	"bufio"
	"fmt"
	"golang-database-project/model"
	"golang-database-project/utils"
	"os"
	"strconv"
	"strings"
)

func main() {
	// services := model.GetAllService()
	//
	// for _, service := range services {
	// 	fmt.Println(service.Id, service.Service_name, service.Satuan, service.Price)
	// }

	// service := model.GetServiceById("S001")
	// fmt.Println(service.Id, service.Service_name, service.Satuan, service.Price)

	// customer := createCustomer()
	// fmt.Println(customer)
	// model.AddCustomer(customer)
	initMenu()
}

func initMenu() {
	for {
		fmt.Println(strings.Repeat("=", 50))
		fmt.Println(strings.Repeat("=", 14), "Enigma Laundry", strings.Repeat("=", 14))
		fmt.Println("1. View All Data")
		fmt.Println("2. Add New Data")
		fmt.Println("3. Update Data")
		fmt.Println("4. Delete Data")
		fmt.Println("5. Exit")
		fmt.Println(strings.Repeat("=", 50))

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter your choice : ")
		scanner.Scan()
		choice, _ := strconv.Atoi(scanner.Text())

		switch choice {
		case 1:
			os.Exit(0)
		case 2:
			choice = utils.AddPrompt()

			switch choice {
			case 1:
				customer := utils.CreateCustomer()
				err := model.AddCustomer(customer)
				if err != nil {
					fmt.Println("Error", err, "\n", "Please Try Again!")
				}
			case 2:
				os.Exit(0)
			case 3:
				os.Exit(0)
			case 4:
				os.Exit(0)
			default:
				fmt.Println("Invalid choice. Please try again.")
			}

		case 3:
			os.Exit(0)
		case 4:
			os.Exit(0)
		case 5:
			fmt.Println("See You Later ^-^")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
