package team

import (
	"errors"
	"strings"

	"github.com/mattismoel/konnekt/internal/cfg"
)

var (
	ErrTeamIDInvalid          = errors.New("Team ID must be a valid positive integer")
	ErrTeamNameInvalid        = errors.New("Team name must be a valid non-empty string")
	ErrTeamDisplayNameInvalid = errors.New("Team display name must be a valid non-empty string")
	ErrTeamDescriptionInvalid = errors.New("Team description must be a valid non-empty string")
)

type Team struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
}

type TeamCollection []Team

func NewTeam(cfgs ...cfg.Func[Team]) (Team, error) {
	t := &Team{}

	err := t.WithCfgs(cfgs...)
	if err != nil {
		return Team{}, err
	}
	return *t, nil
}

func (t *Team) WithCfgs(cfgs ...cfg.Func[Team]) error {
	for _, cfg := range cfgs {
		if err := cfg(t); err != nil {
			return err
		}
	}

	return nil
}

func WithID(id int64) cfg.Func[Team] {
	return func(t *Team) error {
		if id <= 0 {
			return ErrTeamIDInvalid
		}

		t.ID = id
		return nil
	}
}

func WithName(name string) cfg.Func[Team] {
	name = strings.TrimSpace(name)
	return func(t *Team) error {
		if name == "" {
			return ErrTeamNameInvalid
		}

		t.Name = name
		return nil
	}
}

func WithDisplayName(displayName string) cfg.Func[Team] {
	displayName = strings.TrimSpace(displayName)
	return func(t *Team) error {
		if displayName == "" {
			return ErrTeamDisplayNameInvalid
		}

		t.DisplayName = displayName
		return nil
	}
}

func WithDescription(desc string) cfg.Func[Team] {
	desc = strings.TrimSpace(desc)
	return func(t *Team) error {
		if desc == "" {
			return ErrTeamDescriptionInvalid
		}

		t.Description = desc
		return nil
	}
}
