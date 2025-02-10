package mongodb

type Model interface {
	CollectionName() string
}

type Facet[T Model] struct {
	Items []T   `json:"items" bson:"items"`
	Total int64 `json:"total" bson:"total"`
}

func (s Facet[T]) GetItems() []T {
	return s.Items
}

func (s Facet[T]) GetTotal() int64 {
	return s.Total
}
