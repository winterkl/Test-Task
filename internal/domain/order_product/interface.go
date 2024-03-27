package order_product

import (
	order_product_entity "TestTask/internal/domain/order_product/entity"
)

type UseCase interface {
	GetOrderProductList(orders []int) ([]order_product_entity.OrderProduct, error)
}

type Repository interface {
	GetOrderProductList(orders []int) ([]order_product_entity.OrderProduct, error)
}
