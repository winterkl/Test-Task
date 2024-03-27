package category_rack_entity

import (
	rack_entity "TestTask/internal/domain/rack/entity"
)

type CategoryRack struct {
	ID         int `bun:"id,pk,autoincrement"`
	CategoryID int
	RackID     int
	Rack       rack_entity.Rack `bun:"rel:belongs-to,join:rack_id=id"`
	IsMain     bool
}
