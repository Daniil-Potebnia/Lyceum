package grpc

import (
	"context"
	"fmt"
	"service/internal/models"

	client "service/pkg/api/order"
)

type Service interface {
	CreatePosition(ctx context.Context, order models.Position) (*models.Position, error)
	GetPosition(ctx context.Context, order models.Position) (*models.Position, error)
	UpdatePosition(ctx context.Context, order models.Position) (*models.Position, error)
	DeletePosition(ctx context.Context, order models.Position) (bool, error)
	ListPositions(ctx context.Context) ([]*models.Position, error)
}

type OrderService struct {
	client.UnimplementedOrderServiceServer
	service Service
}

func NewOrderService(srv Service) *OrderService {
	return &OrderService{service: srv}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *client.CreateOrderRequest) (*client.CreateOrderResponse, error) {
	if req.Quantity > 0 {
		ord, err := s.service.CreatePosition(ctx, models.Position{Item: req.Item, Quantity: req.Quantity})
		return &client.CreateOrderResponse{Id: ord.ID}, err
	}
	ord, err := &models.Position{}, fmt.Errorf("quantity cannot be less than 1")
	return &client.CreateOrderResponse{Id: ord.ID}, err
}

func (s *OrderService) GetOrder(ctx context.Context, req *client.GetOrderRequest) (*client.GetOrderResponse, error) {
	ord, err := s.service.GetPosition(ctx, models.Position{ID: req.Id})
	return &client.GetOrderResponse{Order: &client.Order{Id: ord.ID, Item: ord.Item, Quantity: ord.Quantity}}, err
}

func (s *OrderService) UpdateOrder(ctx context.Context, req *client.UpdateOrderRequest) (*client.UpdateOrderResponse, error) {
	if req.Quantity > 0 {
		ord, err := s.service.UpdatePosition(ctx, models.Position{ID: req.Id, Item: req.Item, Quantity: req.Quantity})
		return &client.UpdateOrderResponse{Order: &client.Order{Id: ord.ID, Item: ord.Item, Quantity: ord.Quantity}}, err
	}
	ord, err := &models.Position{}, fmt.Errorf("quantity cannot be less than 1")
	return &client.UpdateOrderResponse{Order: &client.Order{Id: ord.ID, Item: ord.Item, Quantity: ord.Quantity}}, err
}

func (s *OrderService) DeleteOrder(ctx context.Context, req *client.DeleteOrderRequest) (*client.DeleteOrderResponse, error) {
	res, err := s.service.DeletePosition(ctx, models.Position{ID: req.Id})
	return &client.DeleteOrderResponse{Success: res}, err
}

func (s *OrderService) ListOrders(ctx context.Context, req *client.ListOrdersRequest) (*client.ListOrdersResponse, error) {
	ords, err := s.service.ListPositions(ctx)
	res := []*client.Order{}
	for _, ord := range ords {
		res = append(res, &client.Order{Id: ord.ID, Item: ord.Item, Quantity: ord.Quantity})
	}
	return &client.ListOrdersResponse{Orders: res}, err
}
