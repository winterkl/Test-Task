package order_entity

import "time"

type Order struct {
	ID        int       `bun:"id,pk,autoincrement"`
	CreatedAt time.Time `bun:",nullzero"`
}
