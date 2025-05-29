package service

import (
	"context"
	"image"
	"io"
	"net/url"
	"path"

	"github.com/mattismoel/konnekt/internal/domain/member"
	"github.com/mattismoel/konnekt/internal/domain/team"
	"github.com/mattismoel/konnekt/internal/object"
	"github.com/mattismoel/konnekt/internal/query"
	"github.com/nfnt/resize"
)

const PROFILE_PICTURE_WIDTH_PX = 512

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

func (srv MemberService) UploadProfilePicture(ctx context.Context, r io.Reader) (string, error) {
	img, _, err := image.Decode(r)
	if err != nil {
		return "", err
	}

	if img.Bounds().Max.X > PROFILE_PICTURE_WIDTH_PX {
		img = resize.Resize(PROFILE_PICTURE_WIDTH_PX, 0, img, resize.Lanczos2)
	}

	formatedImg, err := formatJPEG(img)
	if err != nil {
		return "", err
	}

	fileName := createRandomImageFileName("jpeg")

	url, err := srv.objectStore.Upload(ctx, path.Join("/members", fileName), formatedImg)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (srv MemberService) DeleteProfilePicture(ctx context.Context, url string) error {
	err := srv.objectStore.Delete(ctx, url)
	if err != nil {
		return err
	}

	return nil
}

func (srv MemberService) Approve(ctx context.Context, memberID int64) error {
	err := srv.memberRepo.Approve(ctx, memberID)
	if err != nil {
		return err
	}

	return nil
}

func (srv MemberService) Delete(ctx context.Context, memberID int64) error {
	m, err := srv.memberRepo.ByID(ctx, memberID)
	if err != nil {
		return err
	}

	url, err := url.Parse(m.ProfilePictureURL)
	if err != nil {
		return err
	}

	err = srv.objectStore.Delete(ctx, url.Path)
	if err != nil {
		return err
	}

	err = srv.memberRepo.Delete(ctx, memberID)
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
