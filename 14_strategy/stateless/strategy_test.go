package stateless

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderService_Discount(t *testing.T) {
	orderService := &OrderService{}
	order := &Order{
		Amount: 100,
		Type:   "normal",
	}
	_, err := orderService.Discount(order)
	assert.Equal(t, nil, err)
}
