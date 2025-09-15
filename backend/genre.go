package konnekt

type Genre string

type CreateGenre struct {
	Name      string `json:"name"`
	CreatedBy int64  `json:"-"`
}
