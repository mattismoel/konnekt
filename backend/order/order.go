package order

const (
	ASCENDING  = "ASC"
	DESCENDING = "DESC"
)

type Order string

type Map map[string]Order
