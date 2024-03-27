package order_product_usecase

import (
	"TestTask/internal/domain/order_product"
	order_product_model "TestTask/internal/domain/order_product/model"
	"fmt"
)

type OrderProductUseCase struct {
	orderProductRepo order_product.Repository
}

func NewOrderProductUseCase(orderProductRepo order_product.Repository) *OrderProductUseCase {
	return &OrderProductUseCase{
		orderProductRepo: orderProductRepo,
	}
}

func (uc *OrderProductUseCase) GetOrderProductList(orders []int) (map[string][]order_product_model.GetOrderProductDTO, error) {
	orderList, err := uc.orderProductRepo.GetOrderProductList(orders)
	if err != nil {
		return nil, fmt.Errorf("OrderProductUseCase - "+
			"GetOrderProductList - orderProductRepo.GetOrderProductList: %w", err)
	}
	return order_product_model.NewGetOrderProductListResponse(orderList), nil
}
