package service

import (
	"context"
	"service/internal/models"
)

type OrderRepo interface {
	CreatePosition(ctx context.Context, order models.Position) (*models.Position, error)
	GetPosition(ctx context.Context, order models.Position) (*models.Position, error)
	UpdatePosition(ctx context.Context, order models.Position) (*models.Position, error)
	DeletePosition(ctx context.Context, order models.Position) (bool, error)
	ListPositions(ctx context.Context) ([]*models.Position, error)
}

type OrderService struct {
	Repo OrderRepo
}

func NewOrderService(repo OrderRepo) *OrderService {
	return &OrderService{Repo: repo}
}

func (s *OrderService) CreatePosition(ctx context.Context, order models.Position) (*models.Position, error) {
	return s.Repo.CreatePosition(ctx, order)
}

func (s *OrderService) GetPosition(ctx context.Context, order models.Position) (*models.Position, error) {
	return s.Repo.GetPosition(ctx, order)
}

func (s *OrderService) UpdatePosition(ctx context.Context, order models.Position) (*models.Position, error) {
	return s.Repo.UpdatePosition(ctx, order)
}

func (s *OrderService) DeletePosition(ctx context.Context, order models.Position) (bool, error) {
	return s.Repo.DeletePosition(ctx, order)
}

func (s *OrderService) ListPositions(ctx context.Context) ([]*models.Position, error) {
	return s.Repo.ListPositions(ctx)
}
