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

func ValidateOrderDetailId(id string) error {
	hasPrefix := strings.HasPrefix(id, "D")
	if !hasPrefix {
		return fmt.Errorf("the first character must be '%s'", "D")
	}
	return nil
}

func ValidateQuantity(quantity string) (int, error) {
	newQuantity, err := strconv.Atoi(quantity)
	if err != nil {
		return 0, err
	}

	return newQuantity, nil
}

func CheckOrderDetailId(id string) error {
	_, err := model.GetOrderDetailById(id)
	if err != nil {
		return errors.New("the order_detail does't exist")
	}
	return nil
}

func ViewOrderDetail() {
	orderDetails := model.GetAllOrderDetail()

	fmt.Println()
	fmt.Println("OrderDetails : ")
	if len(orderDetails) < 1 {
		fmt.Println("No OrderDetail  data (empty table)")
	} else {
		for _, orderDetail := range orderDetails {
			fmt.Printf("%s %s %s %d \n", orderDetail.Order_id, orderDetail.Order_id,
				orderDetail.Service_id, orderDetail.Quantity)
		}
	}
	fmt.Println()
}

func ViewAllOrderDetailId() {
	orderDetails := model.GetAllOrderDetail()

	fmt.Println("AvailableOrderDetail: ")
	for _, orderDetail := range orderDetails {
		fmt.Printf("%s \n", orderDetail.Id)
	}
}

func CreateOrderDetail() model.OrderDetail {
	var newOrderDetail model.OrderDetail

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(strings.Repeat("=", 14), "Add New Order Detail", strings.Repeat("=", 14))
	fmt.Println("Enter order Details")

	fmt.Println("Example format for Id 'D001'")
	for {
		fmt.Print("Order Detail Id : ")
		scanner.Scan()
		newOrderDetail.Id = scanner.Text()
		err := ValidateOrderDetailId(newOrderDetail.Id)
		if err != nil {
			fmt.Println(err)
		} else {
			break
		}

	}

	fmt.Println()
	fmt.Print("Example format for order_id 'O001'\n")
	for {
		fmt.Print("OrderDetail order_id : ")
		scanner.Scan()
		newOrderDetail.Order_id = scanner.Text()
		err := ValidateOrderId(newOrderDetail.Order_id)
		err2 := CheckOrderId(newOrderDetail.Order_id)
		if err != nil {
			fmt.Println(err)
		} else if err2 != nil {
			fmt.Println(err2)
		} else {
			break
		}
	}
	fmt.Println()
	fmt.Print("Example format for Service_id 'S001'\n")
	for {
		fmt.Print("OrderDetail Service_id  : ")
		scanner.Scan()
		newOrderDetail.Service_id = scanner.Text()
		err := ValidateServiceId(newOrderDetail.Service_id)
		err2 := CheckServiceId(newOrderDetail.Service_id)
		if err != nil {
			fmt.Println(err)
		} else if err2 != nil {
			fmt.Println(err2)
		} else {
			break
		}
	}

	for {
		fmt.Print("Quantity : ")
		scanner.Scan()
		quantity, err := ValidateQuantity(scanner.Text())
		if err != nil {
			fmt.Println("Quantity must be consist of numbers!")
		} else {
			newOrderDetail.Quantity = quantity
			break
		}

	}

	return newOrderDetail
}
