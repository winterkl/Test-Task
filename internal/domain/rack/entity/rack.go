package rack_entity

type Rack struct {
	ID    int `bun:"id,pk,autoincrement"`
	Title string
}
