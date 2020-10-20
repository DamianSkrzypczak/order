package order

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderNamesListing(t *testing.T) {
	orderfile := Orderfile{
		Version: "0.0.0",
		Orders: map[string]Order{
			"testOrder": {
				Script: []Cmd{"testCmd"},
			},
		},
	}

	assert.Equal(t, []string{"testOrder"}, orderfile.ListOrdersNames())
}

func TestOrderNamesListingWhenNoOrders(t *testing.T) {
	orderfile := Orderfile{
		Version: "0.0.0",
		Orders:  map[string]Order{},
	}

	assert.Equal(t, []string{}, orderfile.ListOrdersNames())
}

func TestOrderfileOrderGetter(t *testing.T) {
	expectedOrder := Order{
		Script: []Cmd{"testCmd"},
	}

	orderfile := Orderfile{
		Version: "0.0.0",
		Orders: map[string]Order{
			"testOrder": expectedOrder,
		},
	}

	order, ok := orderfile.GetOrder("testOrder")
	assert.True(t, ok)
	assert.Equal(t, &expectedOrder, order)
}

func TestOrderfileOrderGetterWenNoOrders(t *testing.T) {
	orderfile := Orderfile{
		Version: "0.0.0",
		Orders:  map[string]Order{},
	}
	_, ok := orderfile.GetOrder("testOrder")
	assert.False(t, ok)
}
