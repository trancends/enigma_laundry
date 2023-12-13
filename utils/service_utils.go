package utils

import (
	"bufio"
	"errors"
	"fmt"
	"golang-database-project/model"
	"os"
	"strconv"
	"strings"
)

func ValidateServiceId(id string) error {
	hasPrefix := strings.HasPrefix(id, "S")
	if !hasPrefix {
		return fmt.Errorf("the first character must be '%s'", "S")
	}
	return nil
}

func ValidateServicePrice(price string) (int, error) {
	newPrice, err := strconv.Atoi(price)
	if err != nil {
		return 0, err
	}

	return newPrice, nil
}

func CheckServiceId(id string) error {
	_, err := model.GetServiceById(id)
	if err != nil {
		return errors.New("service does't exist")
	}
	return nil
}

func ViewService() {
	services := model.GetAllService()

	fmt.Println()
	fmt.Println("Services : ")
	for _, service := range services {
		fmt.Printf("%s, %s, %s, %v\n", service.Id, service.Service_name, service.Satuan, service.Price)
	}
	fmt.Println()
}

func ViewAllServiceId() {
	services := model.GetAllService()

	fmt.Println("AvailableServices: ")
	for _, service := range services {
		fmt.Printf("%s \n", service.Id)
	}
}

func DeleteServiceUtil() {
	var id string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Please enter service id to be deleted : ")
	scanner.Scan()
	id = scanner.Text()
	err := CheckServiceId(id)
	if err != nil {
		fmt.Println(err)
	} else {
		model.DeleteService(id)
	}
}

func CreateService() model.Service {
	var newService model.Service

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(strings.Repeat("=", 14), "Add New service", strings.Repeat("=", 14))
	fmt.Println("Enter Service Details")

	fmt.Println("Example format for Id 'S001'")
	for {
		fmt.Print("Service Id : ")
		scanner.Scan()
		newService.Id = scanner.Text()
		err := ValidateServiceId(newService.Id)
		if err != nil {
			fmt.Println(err)
		} else {
			break
		}

	}

	fmt.Print("Service Name : ")
	scanner.Scan()
	newService.Service_name = scanner.Text()

	fmt.Print("Service Satuan : ")
	scanner.Scan()
	newService.Satuan = scanner.Text()

	for {

		fmt.Print("Service Price : ")
		scanner.Scan()
		price := scanner.Text()
		newPrice, err := ValidateServicePrice(price)
		if err != nil {
			fmt.Println("Price must be consist of numbers!")
		} else {

			newService.Price = newPrice
			break
		}
	}

	return newService
}
