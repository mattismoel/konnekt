package mask

type FieldName string
type FieldValue any
type FieldMap map[FieldName]FieldValue

type Fielder interface {
	Fields() FieldMap
}
