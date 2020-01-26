package model

import "github.com/google/uuid"

type OrderItem struct {
	MenuNumber  int
	Description string
	IsDrink     bool
	Price       float32
}

type DrinksOrder struct {
	Guid  uuid.UUID
	Items []OrderItem
}

type FoodOrder struct {
	Guid  uuid.UUID
	Items []OrderItem
}

type PlaceOrder struct {
	Guid  uuid.UUID
	Items []OrderItem
}
