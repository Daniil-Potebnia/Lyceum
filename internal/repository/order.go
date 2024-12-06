package repository

import (
	"context"
	"fmt"
	"service/internal/models"
	"service/pkg/db/postgres"
)

type OrderRepository struct {
	db *postgres.DB
}

func NewOrderRepository(db *postgres.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (s *OrderRepository) CreatePosition(ctx context.Context, order models.Position) (*models.Position, error) {
	var res models.Position
	err := s.db.Db.QueryRow("INSERT INTO orders (item, quantity) VALUES ($1, $2) RETURNING *", order.Item, order.Quantity).Scan(&res.ID, &res.Item, &res.Quantity)
	if err != nil {
		return &res, fmt.Errorf("error in CreatePosition: %v", err)
	}
	return &res, nil
}

func (s *OrderRepository) GetPosition(ctx context.Context, order models.Position) (*models.Position, error) {
	var res models.Position
	err := s.db.Db.QueryRow("SELECT * FROM orders WHERE id = $1", order.ID).Scan(&res.ID, &res.Item, &res.Quantity)
	if err != nil {
		return &res, fmt.Errorf("error in GetPosition: %v", err)
	}
	return &res, nil
}

func (s *OrderRepository) UpdatePosition(ctx context.Context, order models.Position) (*models.Position, error) {
	var res models.Position
	err := s.db.Db.QueryRow("UPDATE orders SET item = $1, quantity = $2 WHERE id = $3 RETURNING *;", order.Item, order.Quantity, order.ID).Scan(&res.ID, &res.Item, &res.Quantity)
	if err != nil {
		return &res, fmt.Errorf("error in UpdatePosition: %v", err)
	}
	return &res, nil
}

func (s *OrderRepository) DeletePosition(ctx context.Context, order models.Position) (bool, error) {
	res, err := s.db.Db.Exec("DELETE FROM orders WHERE id = $1 RETURNING *;", order.ID)
	if err != nil {
		return false, fmt.Errorf("error in DeletePosition: %v", err)
	}
	num, err := res.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("error in DeletePosition: %v", err)
	}
	if num == 0 {
		return false, fmt.Errorf("error in DeletePosition: %v", fmt.Errorf("there is no order with this id"))
	}
	return true, nil
}

func (s *OrderRepository) ListPositions(ctx context.Context) ([]*models.Position, error) {
	var res []*models.Position
	rows, err := s.db.Db.Query("SELECT * FROM orders;")
	if err != nil {
		return []*models.Position{}, fmt.Errorf("error in ListPositions: %v", err)
	}
	for rows.Next() {
		pos := &models.Position{}
		rows.Scan(&pos.ID, &pos.Item, &pos.Quantity)
		res = append(res, pos)
	}
	return res, nil
}
