package content

import "time"

type LandingImage struct {
	ID  int64  `json:"id"`
	URL string `json:"url"`

	CreatedAt time.Time `json:"createdAt"`
	CreatedBy int64     `json:"createdBy"`
}
