package laba7

import (
	"fmt"
)

type Car struct {
	price float64
	name  string
	color string
	model string
	brand string
}

func (c *Car) getPrice() float64 {
	return c.price
}

func (c *Car) setPrice(newPrice float64) {
	c.price = newPrice
}

func (c *Car) applyDiscount(discount float64) {
	c.price = (100 - discount) * c.price / 100
}

func (c *Car) getModel() string {
	return c.model
}

func (c *Car) getName() string {
	return c.name
}

func (c *Car) getBrand() string {
	return c.brand
}

func (c *Car) getColor() string {
	return c.color
}

func (c *Car) setColor(newName string) {
	c.color = newName
}

func (c *Car) ProductInfo() string {
	return fmt.Sprintf("Машина: %s, Бренд: %s, Модель машины: %s, Цвет машины: %s, Цена машины: %.2f", c.name, c.brand, c.model, c.color, c.price)
}
