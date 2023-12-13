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
			choice = utils.ViewPrompt()

			switch choice {
			case 1:
				utils.ViewCustomer()
			case 2:
				utils.ViewService()
			case 3:
				utils.ViewOrder()
			case 4:
				os.Exit(0)
			default:
				fmt.Println("Invalid choice. Please try again.")
				os.Exit(0)
			}

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
				service := utils.CreateService()
				err := model.AddService(service)
				if err != nil {
					fmt.Println("Error", err, "\n", "Please Try Again!")
				}
			case 3:
				os.Exit(0)
			case 4:
				os.Exit(0)
			default:
				fmt.Println("Invalid choice. Please try again.")
			}

		case 3:
			choice = utils.UpdatePrompt()
			switch choice {
			case 1:
				fmt.Println()
				utils.ViewAllCustomerId()
				fmt.Println()
				fmt.Println("Enter Detail Below to Updates : ")
				customer := utils.CreateCustomer()
				err := model.UpdateCustomer(customer)
				if err != nil {
					fmt.Println("Error", "Customer doesn't Exist!", "\n", "Please Try Again!")
				}
			case 2:
				fmt.Println()
				utils.ViewAllServiceId()
				fmt.Println()
				fmt.Println("Enter Detail Below to Updates : ")
				service := utils.CreateService()
				err := model.UpdateService(service)
				if err != nil {
					fmt.Println("Error", "Service doesn't Exist!", "\n", "Please Try Again!")
				}
			case 3:
				os.Exit(0)
			case 4:
				os.Exit(0)
			default:
				fmt.Println("Invalid choice. Please try again.")
			}
		case 4:
			choice = utils.DeletePrompt()

			switch choice {
			case 1:
				utils.ViewAllCustomerId()
				utils.DeleteCustomerUtil()
			case 2:
				os.Exit(0)
			default:
				fmt.Println("Invalid choice. Please try again.")
			}
		case 5:
			fmt.Println("See You Later ^-^")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
