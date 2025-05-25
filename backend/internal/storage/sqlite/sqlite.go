package sqlite

type Scanner interface {
	Scan(dst ...any) error
}
