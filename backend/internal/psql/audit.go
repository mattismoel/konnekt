package psql

import "time"

type Timestamps struct {
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type AuditFields struct {
	CreatedBy int64 `db:"created_by"`
	UpdatedBy int64 `db:"updated_by"`
}
