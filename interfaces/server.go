package interfaces

import (
	"ms-sample/order/interfaces/service"

	pb "ms-sample/order/proto/order"

	"ms-sample/order/app/usecase/order"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"ms-sample/order/domain/repository"
)

type ServerParams struct {
	OrderRepository repository.OrderRepository
	EventRepository repository.EventRepository
}

func NewServer(params ServerParams) *grpc.Server {
	server := grpc.NewServer()

	orderService := service.NewOrderServer(
		order.NewListOrders(params.OrderRepository),
		order.NewGetOrder(params.OrderRepository),
		order.NewCreateOrder(params.OrderRepository, params.EventRepository),
	)

	reflection.Register(server)

	pb.RegisterOrderServiceServer(server, orderService)

	return server

}
