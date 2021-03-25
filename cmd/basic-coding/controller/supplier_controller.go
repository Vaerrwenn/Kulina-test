package controller

import (
	"bufio"
	"fmt"
	"kulina-test/cmd/basic-coding/entity"
	"os"
	"strings"
)

// SupplierRegistration handles Supplier Registration.
func SupplierRegistration() entity.Supplier {
	var supplier entity.Supplier

	scanner := bufio.NewScanner(os.Stdin)
	for isRunning := true; isRunning; {
		fmt.Print("Insert Supplier Name: ")
		scanner.Scan()
		supplier.Name = scanner.Text()

		if len(supplier.Name) < 4 {
			fmt.Println("Name must have more than 3 characters.")
		} else {
			isRunning = false
		}
	}

	return supplier
}

// SupplierCreateStore handles Store creation.
func SupplierCreateStore(supplier *entity.Supplier) {
	scanner := bufio.NewScanner(os.Stdin)

	for isRunning := true; isRunning; {
		fmt.Print("Insert Store Name: ")
		scanner.Scan()
		supplier.StoreName = scanner.Text()

		if len(supplier.StoreName) < 4 {
			fmt.Println("Name must have more than 3 characters.")
		} else {
			isRunning = false
		}
	}

}

// SupplierCreateAreas handles Selling Areas creation for Supplier.
func SupplierCreateAreas(supplier *entity.Supplier) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Insert Selling Area(s) (If you serve multiple areas, type the areas and separate them with space): ")
	scanner.Scan()

	areasString := scanner.Text()
	areas := strings.Split(areasString, " ")
	supplier.SellingAreas = areas
}
