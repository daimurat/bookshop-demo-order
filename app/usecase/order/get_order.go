package order

import (
	"context"
	"ms-sample/order/domain/model"
	"ms-sample/order/domain/repository"
)

type GetOrdersParams struct {
	ID string
}

type GetOrder func(ctx context.Context, params GetOrdersParams) (*model.Order, error)

func NewGetOrder(orderRepository repository.OrderRepository) GetOrder {
	return func(ctx context.Context, params GetOrdersParams) (*model.Order, error) {
		return orderRepository.GetOrder(ctx, params.ID)
	}
}
