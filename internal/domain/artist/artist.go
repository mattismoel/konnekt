package artist

type Artist struct {
	ID          int64
	Name        string
	Description string
	ImageURL    string
	Genres      []Genre
	Socials     []Social
}
