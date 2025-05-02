package repository

import (
	"context"

	"ms-sample/order/domain/model"
)

type EventRepository interface {
	PostOrderEvent(ctx context.Context, orderEvent model.OrderEvent) error
}
