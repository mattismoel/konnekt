package artist

type Artist struct {
	ID          int64    `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	ImageURL    string   `json:"imageUrl"`
	Genres      []Genre  `json:"genres"`
	Socials     []Social `json:"socials"`
}

func NewArtist(name, description, imageURL string, socials []Social, genres []Genre) (Artist, error) {
	return Artist{
		Name:        name,
		Description: description,
		ImageURL:    imageURL,
		Socials:     socials,
		Genres:      genres,
	}, nil
}

func (a *Artist) WithGenres(genres []Genre) {
	a.Genres = genres
}

func (a *Artist) WithSocials(socials []Social) {
	a.Socials = socials
}
