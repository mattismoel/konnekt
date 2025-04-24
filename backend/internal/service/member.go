package service

import (
	"context"
	"fmt"
	"io"
	"path"

	"github.com/google/uuid"
	"github.com/mattismoel/konnekt/internal/domain/member"
	"github.com/mattismoel/konnekt/internal/domain/team"
	"github.com/mattismoel/konnekt/internal/object"
	"github.com/mattismoel/konnekt/internal/query"
)

type MemberService struct {
	memberRepo  member.Repository
	teamRepo    team.Repository
	objectStore object.Store
}

func NewMemberService(memberRepo member.Repository, teamRepo team.Repository, objectStore object.Store) (*MemberService, error) {
	return &MemberService{
		memberRepo:  memberRepo,
		teamRepo:    teamRepo,
		objectStore: objectStore,
	}, nil
}

func (srv MemberService) ByID(ctx context.Context, memberID int64) (member.Member, error) {
	u, err := srv.memberRepo.ByID(ctx, memberID)
	if err != nil {
		return member.Member{}, err
	}

	return u, nil
}

func (srv MemberService) List(ctx context.Context, q query.ListQuery) (query.ListResult[member.Member], error) {
	result, err := srv.memberRepo.List(ctx, q)
	if err != nil {
		return query.ListResult[member.Member]{}, err
	}

	return result, nil
}

func (srv MemberService) UploadProfilePicture(ctx context.Context, fileName string, r io.Reader) (string, error) {
	ext := path.Ext(fileName)

	fileName = fmt.Sprintf("%s%s", uuid.NewString(), ext)

	url, err := srv.objectStore.Upload(ctx, path.Join("/members", fileName), r)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (srv MemberService) Approve(ctx context.Context, memberID int64) error {
	err := srv.memberRepo.Approve(ctx, memberID)
	if err != nil {
		return err
	}

	return nil
}

func (srv MemberService) Delete(ctx context.Context, memberID int64) error {
	err := srv.memberRepo.Delete(ctx, memberID)
	if err != nil {
		return err
	}
	return nil
}

func (srv MemberService) Update(ctx context.Context, memberID int64, m member.Member) error {
	if err := srv.memberRepo.Update(ctx, memberID, m); err != nil {
		return err
	}

	return nil
}

func (srv MemberService) SetMemberTeams(ctx context.Context, memberID int64, teamIDs ...int64) error {
	teams := make(team.TeamCollection, 0)

	// Check that all teams exist. If not return.
	for _, teamID := range teamIDs {
		team, err := srv.teamRepo.ByID(ctx, teamID)
		if err != nil {
			return err
		}

		teams = append(teams, team)
	}

	err := srv.memberRepo.SetMemberTeams(ctx, memberID, teamIDs...)
	if err != nil {
		return err
	}

	return nil
}
