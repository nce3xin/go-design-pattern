package stateful

import "fmt"

type IDiscountStrategy interface {
	calDiscount(order *Order) float64
}

// NormalDiscountStrategy 普通订单折扣
type NormalDiscountStrategy struct {
}

func (n NormalDiscountStrategy) calDiscount(order *Order) float64 {
	//TODO implement me
	return 0
}

// GroupDiscountStrategy 团购订单折扣
type GroupDiscountStrategy struct {
}

func (g GroupDiscountStrategy) calDiscount(order *Order) float64 {
	//TODO implement me
	return 0
}

// VIPDiscountStrategy VIP订单折扣
type VIPDiscountStrategy struct {
}

func (V VIPDiscountStrategy) calDiscount(order *Order) float64 {
	//TODO implement me
	return 0
}

// Order 订单
type Order struct {
	Amount float64
	Type   string
}

// GetStrategyInstance 如果策略类是有状态的，需要每次都创建新的实例。
// 本来应该创建一个工厂类，根据不同的策略类型，创建不同的策略实例。
// 但由于go不支持类的静态方法，推荐使用一个全局的函数来创建策略实例。
func GetStrategyInstance(orderType string) (IDiscountStrategy, error) {
	switch orderType {
	case "normal":
		return NormalDiscountStrategy{}, nil
	case "group":
		return GroupDiscountStrategy{}, nil
	case "vip":
		return VIPDiscountStrategy{}, nil
	default:
		return nil, fmt.Errorf("not found strategy for type %s", orderType)
	}
}

// OrderService 订单服务，策略的使用
type OrderService struct {
}

func (os *OrderService) Discount(order *Order) (float64, error) {
	orderType := order.Type
	discountStrategy, err := GetStrategyInstance(orderType)
	if err != nil {
		return 0, err
	}
	discount := discountStrategy.calDiscount(order)
	return discount, nil
}
