package product_entity

import (
	product_categories_entity "TestTask/internal/domain/product_categories/entity"
)

type Product struct {
	ID         int `bun:"id,pk,autoincrement"`
	Title      string
	CategoryID int
	Category   product_categories_entity.ProductCategory `bun:"rel:belongs-to,join:category_id=id"`
}
