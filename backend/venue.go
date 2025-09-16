package konnekt

import "github.com/mattismoel/konnekt/backend/mask"

type Venue struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	City    string `json:"city"`
	Country string `json:"country"`
}

type CreateVenue struct {
	Name      string `json:"name"`
	City      string `json:"city"`
	Country   string `json:"country"`
	CreatedBy int64  `json:"-"`
}

type UpdateVenue struct {
	Name      string `json:"name"`
	City      string `json:"city"`
	Country   string `json:"country"`
	UpdatedBy int64  `json:"-"`
}

func (uv UpdateVenue) Fields() mask.FieldMap {
	return mask.FieldMap{
		"name":       uv.Name,
		"city":       uv.City,
		"country":    uv.Country,
		"updated_by": uv.UpdatedBy,
	}
}
