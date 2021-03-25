package entity

type Supplier struct {
	Name         string
	StoreName    string
	Product      Product
	SellingAreas []string // For example, A B C D E
}
