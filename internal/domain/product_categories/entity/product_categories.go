package product_categories_entity

import category_rack_entity "TestTask/internal/domain/category_rack/entity"

type ProductCategory struct {
	ID           int `bun:"id,pk,autoincrement"`
	Title        string
	CategoryRack []category_rack_entity.CategoryRack `bun:"rel:has-many,join:id=category_id"`
}
