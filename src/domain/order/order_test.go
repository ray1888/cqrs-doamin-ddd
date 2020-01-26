package order

import (
	"testing"

	"github.com/cqrs-domain-ddd/v1/src/domain/order/model"
	"github.com/google/uuid"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestCanOrderWithUnopenedTab(t *testing.T) {
	Convey("Given Nothing", t, func() {
		Convey("When PlaceOrder", func() {
			UUIDOrder, _ := uuid.NewUUID()
			po := model.PlaceOrder{UUIDOrder, nil}
			err := PlaceOrderHandlerImpl{}.Handle(po)
			assert.EqualError(t, err, TabNotOpenError.Error())
		})
	})
}

func TestCanPlaceFoodOrder(t *testing.T) {

}

func TestCanPlaceDrinksOrder(t *testing.T) {

}

func TestCanPlaceFoodAndDrinkOrder(t *testing.T) {

}
