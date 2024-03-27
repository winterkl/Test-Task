package order_product_entity

import (
	order_entity "TestTask/internal/domain/order/entity"
	product_entity "TestTask/internal/domain/product/entity"
	rack_entity "TestTask/internal/domain/rack/entity"
)

type OrderProduct struct {
	ID        int `bun:"id,pk,autoincrement"`
	OrderID   int
	Order     order_entity.Order `bun:"rel:belongs-to,join:order_id=id"`
	ProductID int
	Product   product_entity.Product `bun:"rel:belongs-to,join:product_id=id"`
	Count     int
	RackID    int
	Rack      rack_entity.Rack `bun:"rel:belongs-to,join:rack_id=id"`
}
