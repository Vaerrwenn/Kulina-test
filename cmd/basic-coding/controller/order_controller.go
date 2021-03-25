package controller

import (
	"bufio"
	"fmt"
	"kulina-test/cmd/basic-coding/entity"
	"log"
	"os"
	"strconv"
)

// OrderCreate handles Order creation.
func OrderCreate(user *entity.User, supplier entity.Supplier, product entity.Product) *entity.Order {
	var order entity.Order

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\nChoose a Supplier: ")
	fmt.Printf("\n1. %s.\n", supplier.Name)
	fmt.Print("Input: ")
	scanner.Scan()
	inputSup := scanner.Text()

	if inputSup != "1" {
		return nil
	}

	isUserAddressCovered := func() bool {
		for _, area := range supplier.SellingAreas {
			if area == user.Address {
				return true
			}
		}
		return false
	}

	if !isUserAddressCovered() {
		fmt.Println("Area is not covered. Sorry :(")
		return nil
	}

	fmt.Println("\n\nChoose a product to buy: ")
	fmt.Printf("1. %s. Price: %s.\n", product.Name, product.Price)
	fmt.Print("Input: ")
	scanner.Scan()
	inputProd := scanner.Text()

	if inputProd != "1" {
		return nil
	}

	fmt.Println("\n\nChoose purchase method: ")
	fmt.Println("1. Subscription")
	fmt.Println("2. One-time Purchase")
	fmt.Print("Input: ")
	scanner.Scan()
	inputMethod, err := strconv.Atoi(scanner.Text())

	if err != nil {
		log.Fatal("ERROR: Invalid input syntax.")
	}

	if inputMethod < 1 || inputMethod > 2 {
		log.Fatal("ERROR: Invalid input number. Choice does not exist.")
	}

	order = entity.Order{
		UserName:       user.Name,
		ProductName:    product.Name,
		PurchaseOption: inputMethod,
		ProductPrice:   product.Price,
	}

	return &order
}
