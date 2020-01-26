package order

import "github.com/cqrs-domain-ddd/v1/src/domain/order/model"

type PlaceOrderHandler interface {
	Handle(order model.PlaceOrder)
}

type PlaceOrderHandlerImpl struct{}

func (ph PlaceOrderHandlerImpl) Handle(p model.PlaceOrder) error {
	return TabNotOpenError
}

type FoodOrderHandler interface {
	Handle(order model.FoodOrder)
}

type FoodOrderHandlerImpl struct{}

func (fh FoodOrderHandlerImpl) Handle(p model.PlaceOrder) error {
	return TabNotOpenError
}

type DrinksOrderHandler interface {
	Handle(order model.DrinksOrder)
}

type DrinksOrderHandlerImpl struct{}

func (dh DrinksOrderHandlerImpl) Handle(p model.PlaceOrder) error {
	return TabNotOpenError
}
