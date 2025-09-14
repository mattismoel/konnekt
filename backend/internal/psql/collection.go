package psql

type Domainer[T any, K any] interface {
	ToDomain() K
}

type Collection[T Domainer[T, K], K any] []T

func (c Collection[T, K]) ToDomain() []K {
	domainItems := make([]K, 0)
	for _, item := range c {
		domainItems = append(domainItems, item.ToDomain())
	}

	return domainItems
}
