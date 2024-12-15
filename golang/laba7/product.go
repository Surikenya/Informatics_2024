package laba7

type Product interface {
	ProductInfo() string
	applyDiscount(discount float64)
	getPrice() float64
}
