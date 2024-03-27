package order_product_repository

import (
	order_product_entity "TestTask/internal/domain/order_product/entity"
	"TestTask/pkg/postgres"
	"context"
	"fmt"
	"github.com/uptrace/bun"
)

type OrderProductRepository struct {
	db *postgres.Postgres
}

func NewOrderProductRepository(db *postgres.Postgres) *OrderProductRepository {
	return &OrderProductRepository{
		db: db,
	}
}

func (r *OrderProductRepository) GetOrderProductList(orders []int) ([]order_product_entity.OrderProduct, error) {
	orderList := []order_product_entity.OrderProduct{}
	if err := r.db.
		NewSelect().
		Model(&orderList).
		Relation("Order").Relation("Product.Category.CategoryRack.Rack").Relation("Rack").
		Where("order_id IN (?)", bun.In(orders)).
		Order("rack.title").
		Scan(context.TODO()); err != nil {
		return nil, fmt.Errorf("OrderProductRepository - GetOrderProductList - NewSelect: %w", err)
	}

	return orderList, nil
}
