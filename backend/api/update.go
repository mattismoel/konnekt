package api

import (
	"fmt"
	"slices"

	"github.com/mattismoel/konnekt/backend/mask"
)

type UpdateRequest[F mask.Fielder] struct {
	Data       F                `json:"data"`
	UpdateMask []mask.FieldName `json:"updateMask"`
}

// Returns a field-name-to-value map.
func (ur UpdateRequest[F]) UpdateMap() mask.FieldMap {
	um := make(mask.FieldMap)

	validFieldMap := ur.Data.Fields()
	fmt.Println("FIELDS", validFieldMap)

	for fieldName, value := range validFieldMap {
		if slices.Contains(ur.UpdateMask, fieldName) {
			um[fieldName] = value
		}
	}

	return um
}
