package service

import (
	"context"

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

func NewTeamService(teamRepo team.Repository, authRepo auth.Repository) *TeamService {
	return &TeamService{
		teamRepo: teamRepo,
		authRepo: authRepo,
	}
}

func (ts TeamService) ByID(ctx context.Context, teamID int64) (team.Team, error) {
	t, err := ts.teamRepo.ByID(ctx, teamID)
	if err != nil {
		return team.Team{}, err
	}

	return t, nil
}

func (ts TeamService) TeamPermissions(ctx context.Context, teamID int64) (auth.PermissionCollection, error) {
	_, err := ts.teamRepo.ByID(ctx, teamID)
	if err != nil {
		return nil, err
	}

	perms, err := ts.authRepo.TeamPermissions(ctx, teamID)
	if err != nil {
		return nil, err
	}

	return perms, nil
}

func (ts TeamService) Delete(ctx context.Context, teamID int64) error {
	err := ts.teamRepo.Delete(ctx, teamID)
	if err != nil {
		return err
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
		return team.Team{}, err
	}

	teamID, err := ts.teamRepo.Insert(ctx, r)
	if err != nil {
		return team.Team{}, err
	}

	t, err := ts.teamRepo.ByID(ctx, teamID)
	if err != nil {
		return team.Team{}, err
	}

	return t, nil
}

func (ts TeamService) List(ctx context.Context, q query.ListQuery) (query.ListResult[team.Team], error) {
	result, err := ts.teamRepo.List(ctx, q)
	if err != nil {
		return query.ListResult[team.Team]{}, err
	}

	return result, nil
}

func (ts TeamService) MemberTeams(ctx context.Context, memberID int64) (team.TeamCollection, error) {
	_, err := ts.memberRepo.ByID(ctx, memberID)
	if err != nil {
		return nil, err
	}

	teams, err := ts.teamRepo.MemberTeams(ctx, memberID)
	if err != nil {
		return nil, err
	}

	return teams, nil
}
