package main

import (
	"fmt"
	"kulina-test/cmd/basic-coding/controller"
	"kulina-test/cmd/basic-coding/entity"
	"log"
	"strings"
)

func main() {
	var user entity.User
	var supplier entity.Supplier
	var product entity.Product

	for isRunning := true; isRunning; {
		fmt.Println("Register as: ")
		fmt.Println("1. User")
		fmt.Println("2. Supplier")
		fmt.Println("3. Exit")
		fmt.Print("Input number: ")

		var regisInput int
		fmt.Scanln(&regisInput)

		if regisInput == 1 {
			if product.Name == "" || supplier.Name == "" {
				fmt.Println("No product or supplier is available. Please make a product first.")
				break
				// I really don't want to use goto here.
			}

			user = controller.UserRegistration()

			controller.UserInputAddress(&user)

			order := controller.OrderCreate(&user, supplier, product)

			if order == nil {
				fmt.Println("Failed to create an Order.")
				break
			}

			purchaseOptDesc := ""
			if order.PurchaseOption == 1 {
				purchaseOptDesc = "Subscription"
			} else {
				purchaseOptDesc = "One-time Purchase"
			}

			fmt.Println("Confirm Order.")
			fmt.Printf("Your name: %s.\nYour order: %s.\nPurchase method: %s\n. Price: %s.\n",
				order.UserName, order.ProductName, purchaseOptDesc, order.ProductPrice)
			fmt.Print("Confirm Order (Y/N)? ")

			var confirm string
			fmt.Scan(&confirm)

			if strings.ToUpper(confirm) == "Y" {
				log.Fatalf("Oops... Something is broken.\nWe're fixing this.\nPlease come later.")
			} else if strings.ToUpper(confirm) == "N" {
				fmt.Print("\nOrder cancelled.")
				break
			} else {
				log.Fatal("ERROR: Invalid input.")
			}

		} else if regisInput == 2 {
			supplier = controller.SupplierRegistration()

			controller.SupplierCreateStore(&supplier)

			product = controller.ProductCreate(&supplier)

			controller.SupplierCreateAreas(&supplier)

		} else if regisInput == 3 {
			isRunning = false
		} else {
			log.Fatal("Invalid input.")
		}
	}

}
