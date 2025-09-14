package psql

import "time"

type Timestamps struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AuditFields struct {
	CreatedBy int64
	UpdatedBy int64
}
