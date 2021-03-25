package controller

import (
	"bufio"
	"fmt"
	"kulina-test/cmd/basic-coding/entity"
	"os"
)

// ProductCreate handles Product creation.
func ProductCreate(supplier *entity.Supplier) entity.Product {
	scanner := bufio.NewScanner(os.Stdin)

	var product entity.Product

	fmt.Print("Insert Product Name: ")
	scanner.Scan()
	product.Name = scanner.Text()

	fmt.Print("Inser Price of that product: ")
	scanner.Scan()
	product.Price = scanner.Text()

	product.SupplierName = supplier.Name

	supplier.Product = product

	return product
}
