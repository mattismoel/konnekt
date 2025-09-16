package konnekt

import "github.com/mattismoel/konnekt/backend/mask"

type Artist struct {
	ID          int64            `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	ImageURL    string           `json:"imageUrl"`
	PreviewURL  string           `json:"previewUrl"`
	Socials     SocialCollection `json:"social"`
	Genres      []Genre          `json:"genres"`
}

type CreateArtist struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	ImageURL    string           `json:"imageUrl"`
	PreviewURL  string           `json:"previewUrl"`
	Socials     SocialCollection `json:"socials"`
	GenreIDs    []int64          `json:"genreIds"`
	CreatedBy   int64            `json:"-"`
}

type UpdateArtist struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	ImageURL    string           `json:"imageUrl"`
	PreviewURL  string           `json:"previewUrl"`
	Socials     SocialCollection `json:"socials"`
	GenreIDs    []int64          `json:"genres"`
}

type Social string
type SocialCollection []Social

func (ua UpdateArtist) Fields() mask.FieldMap {
	return mask.FieldMap{
		"name":        ua.Name,
		"description": ua.Description,
		"image_url":   ua.ImageURL,
		"preview_url": ua.PreviewURL,
		"socials":     ua.Socials,
		"genres":      ua.GenreIDs,
	}
}

func (sc SocialCollection) String() []string {
	strs := make([]string, 0)
	for _, s := range sc {
		strs = append(strs, string(s))
	}
	return strs
}
