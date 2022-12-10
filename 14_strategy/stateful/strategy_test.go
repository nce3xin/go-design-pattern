package stateful

import "testing"

func TestOrderService_Discount(t *testing.T) {
	orderService := &OrderService{}
	order := &Order{
		Amount: 100,
		Type:   "normal",
	}
	_, err := orderService.Discount(order)
	if err != nil {
		t.Fatalf("err should be nil")
	}
}
