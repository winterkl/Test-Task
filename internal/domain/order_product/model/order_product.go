package order_product_model

import (
	order_product_entity "TestTask/internal/domain/order_product/entity"
	"fmt"
	"strings"
)

type GetOrderProductDTO struct {
	ID             int
	Product        string
	ProductID      int
	Count          int
	MainRack       string
	AdditionalRack []string
}

func (o GetOrderProductDTO) String() string {
	str := fmt.Sprintf("%v (id=%v)\n"+
		"заказ %v, %v шт\n", o.Product, o.ProductID, o.ID, o.Count)
	if o.AdditionalRack != nil {
		str += "доп стеллаж: " + strings.Join(o.AdditionalRack, ", ") + "\n"
	}
	return str
}

func NewGetOrderProductListResponse(orders []order_product_entity.OrderProduct) map[string][]GetOrderProductDTO {
	rackOrderMap := make(map[string][]GetOrderProductDTO)
	for _, vl := range orders {
		orderProduct := GetOrderProductDTO{
			ID:        vl.OrderID,
			Product:   vl.Product.Title,
			ProductID: vl.ProductID,
			Count:     vl.Count,
		}
		for _, rack := range vl.Product.Category.CategoryRack {
			if rack.IsMain == true {
				orderProduct.MainRack = rack.Rack.Title
			} else {
				orderProduct.AdditionalRack = append(orderProduct.AdditionalRack, rack.Rack.Title)
			}
		}
		rackOrderMap[vl.Rack.Title] = append(rackOrderMap[vl.Rack.Title], orderProduct)
	}
	return rackOrderMap
}
