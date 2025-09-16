package konnekt

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
