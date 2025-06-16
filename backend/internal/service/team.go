package service

import (
	"context"
	"fmt"

	"github.com/mattismoel/konnekt/internal/domain/auth"
	"github.com/mattismoel/konnekt/internal/domain/member"
	"github.com/mattismoel/konnekt/internal/domain/team"
	"github.com/mattismoel/konnekt/internal/query"
)

type TeamService struct {
	teamRepo   team.Repository
	memberRepo member.Repository
	authRepo   auth.Repository
}

func NewTeamService(teamRepo team.Repository, memberRepo member.Repository, authRepo auth.Repository) *TeamService {
	return &TeamService{
		teamRepo:   teamRepo,
		memberRepo: memberRepo,
		authRepo:   authRepo,
	}
}

func (ts TeamService) ByID(ctx context.Context, teamID int64) (team.Team, error) {
	t, err := ts.teamRepo.ByID(ctx, teamID)
	if err != nil {
		return team.Team{}, fmt.Errorf("Could not find team %d: %v", teamID, err)
	}

	return t, nil
}

func (ts TeamService) TeamPermissions(ctx context.Context, teamID int64) (auth.PermissionCollection, error) {
	_, err := ts.teamRepo.ByID(ctx, teamID)
	if err != nil {
		return nil, fmt.Errorf("Could not find team %d: %v", teamID, err)
	}

	perms, err := ts.authRepo.TeamPermissions(ctx, teamID)
	if err != nil {
		return nil, fmt.Errorf("Could not get team permissions: %v", err)
	}

	return perms, nil
}

func (ts TeamService) Delete(ctx context.Context, teamID int64) error {
	err := ts.teamRepo.Delete(ctx, teamID)
	if err != nil {
		return fmt.Errorf("Could not delete team with id %d: %v", teamID, err)
	}

	return nil
}

type CreateTeam struct {
	Name        string
	DisplayName string
	Description string
}

func (ts TeamService) Create(ctx context.Context, load CreateTeam) (team.Team, error) {
	r, err := team.NewTeam(
		team.WithName(load.Name),
		team.WithDisplayName(load.DisplayName),
		team.WithDescription(load.Description),
	)

	if err != nil {
		return team.Team{}, fmt.Errorf("Could not create team: %v", err)
	}

	teamID, err := ts.teamRepo.Insert(ctx, r)
	if err != nil {
		return team.Team{}, fmt.Errorf("Could not insert team into repository: %v", err)
	}

	t, err := ts.teamRepo.ByID(ctx, teamID)
	if err != nil {
		return team.Team{}, fmt.Errorf("Could not get team with id %d: %v", teamID, err)
	}

	return t, nil
}

func (ts TeamService) List(ctx context.Context, q query.ListQuery) (query.ListResult[team.Team], error) {
	result, err := ts.teamRepo.List(ctx, q)
	if err != nil {
		return query.ListResult[team.Team]{}, fmt.Errorf("Could not list teams: %v", err)
	}

	return result, nil
}

func (ts TeamService) MemberTeams(ctx context.Context, memberID int64) (team.TeamCollection, error) {
	_, err := ts.memberRepo.ByID(ctx, memberID)
	if err != nil {
		return nil, fmt.Errorf("Could not get member with id %d: %v", memberID, err)
	}

	teams, err := ts.teamRepo.MemberTeams(ctx, memberID)
	if err != nil {
		return nil, fmt.Errorf("Could not get member teams: %v", err)
	}

	return teams, nil
}
