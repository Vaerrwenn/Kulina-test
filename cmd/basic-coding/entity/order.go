package entity

type Order struct {
	UserName       string
	ProductName    string
	PurchaseOption int // 1 = Subscription, 2 = One-time purchase
	ProductPrice   string
}
