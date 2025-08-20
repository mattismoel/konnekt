package sqlite

type AuditFields struct {
	CreatedBy int64
	UpdatedBy int64
}

func (af1 AuditFields) Equals(af2 AuditFields) bool {
	return af1.CreatedBy == af2.CreatedBy && af1.UpdatedBy == af2.UpdatedBy
}
