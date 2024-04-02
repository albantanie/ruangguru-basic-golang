package main

import (
	"a21hc3NpZ25tZW50/database"
	"a21hc3NpZ25tZW50/service"
	"fmt"
)

func CashierApp(db *database.Database) service.ServiceInterface {
	service := service.NewService(db)
	return service
}

func main() {
	database := database.NewDatabase()
	service := CashierApp(database)

	// Test Case AddCart
	err := service.AddCart("Kaos Polos", 2)
	if err != nil {
		panic(err)
	}
	fmt.Println("Test Case AddCart passed")

	// Test Case ShowCart
	cartItems, err := service.ShowCart()
	if err != nil {
		panic(err)
	}
	fmt.Println("Test Case ShowCart passed")
	fmt.Println("Cart Items:", cartItems)

	// Test Case RemoveCart
	err = service.RemoveCart("Kaos Polos")
	if err != nil {
		panic(err)
	}
	fmt.Println("Test Case RemoveCart passed")

	// Test Case ResetCart
	err = service.ResetCart()
	if err != nil {
		panic(err)
	}
	fmt.Println("Test Case ResetCart passed")

	// Test Case GetAllProduct
	products, err := service.GetAllProduct()
	if err != nil {
		panic(err)
	}
	fmt.Println("Test Case GetAllProduct passed")
	fmt.Println("Products:", products)

	// Test Case Pay
	err = service.AddCart("Kaos Polos", 2)
	if err != nil {
		panic(err)
	}
	err = service.AddCart("Kaos Sablon", 1)
	if err != nil {
		panic(err)
	}
	paymentInfo, err := service.Pay(500000)
	if err != nil {
		panic(err)
	}
	fmt.Println("Test Case Pay passed")
	fmt.Println("Payment Information:", paymentInfo)
}
