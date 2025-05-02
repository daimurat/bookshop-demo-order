package order

import (
	"context"
	"ms-sample/order/domain/model"
	"ms-sample/order/domain/repository"
)

type ListOrdersParams struct {
	CustomerId string
}

type ListOrders func(ctx context.Context, params ListOrdersParams) ([]*model.Order, error)

func NewListOrders(orderRepository repository.OrderRepository) ListOrders {
	return func(ctx context.Context, params ListOrdersParams) ([]*model.Order, error) {
		return orderRepository.ListOrders(ctx, params.CustomerId)
	}
}
