package laba7

import (
	"fmt"
)

func CalculationSumProduct(listproducts []Product) string {
	var sum float64 = 0
	for _, product := range listproducts {
		sum += product.getPrice()
	}
	s := fmt.Sprintf("%.2f", sum)
	return s
}

func RunLab7() {
	bubleTea := &Food{name: "Баблти", price: 149.99, weight: 10}
	audi := &Car{price: 1500000, name: "Ауди", color: "Абсолютный черный", model: "А4", brand: "Audi"}
	phone := &Technic{price: 2000000, name: "Audi", color: "черный"}

	bubleTea.setName("Бобати")
	bubleTea.setPrice(159.99)
	bubleTea.getName()
	phone.getName()
	phone.getPrice()
	phone.getColor()
	phone.getModel()
	phone.setPrice(14000)
	phone.getCode()
	phone.getBrand()
	audi.getName()
	audi.setPrice(1600000)
	audi.setColor("белый")
	audi.getColor()
	audi.getBrand()
	audi.getModel()

	listproducts := []Product{bubleTea, audi, phone}
	fmt.Printf("Сумма товаров, без учёта скидки, равна: %v рублей \n", CalculationSumProduct(listproducts))
	bubleTea.applyDiscount(9)
	audi.applyDiscount(23)
	phone.applyDiscount(8)
	fmt.Printf("Сумма товаров, с учётом скидки, равна: %v рублей \n", CalculationSumProduct(listproducts))
}
