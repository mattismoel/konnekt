package konnekt

type Member struct {
	ID          int64  `json:"id"`
	Email       string `json:"email"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	AvatarURL   string `json:"avatarUrl"`
	SpecialRole string `json:"specialRole,omitempty"`
	Teams       []Team `json:"teams"`
	Approved    bool   `json:"approved"`
}

type CreateMember struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	AvatarURL string `json:"avatarUrl"`
	Password  string `json:"password"`
}
