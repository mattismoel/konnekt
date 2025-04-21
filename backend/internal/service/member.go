package service

import (
	"context"
	"fmt"
	"io"
	"path"

	"github.com/google/uuid"
	"github.com/mattismoel/konnekt/internal/domain/member"
	"github.com/mattismoel/konnekt/internal/object"
	"github.com/mattismoel/konnekt/internal/query"
)

type MemberService struct {
	memberRepo  member.Repository
	objectStore object.Store
}

func NewMemberService(memberRepo member.Repository, objectStore object.Store) (*MemberService, error) {
	return &MemberService{
		memberRepo:  memberRepo,
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

func (srv MemberService) SetMemberProfilePicture(ctx context.Context, memberID int64, fileName string, r io.Reader) (string, error) {
	ext := path.Ext(fileName)

	fileName = fmt.Sprintf("%s%s", uuid.NewString(), ext)

	url, err := srv.objectStore.Upload(ctx, path.Join("/members", fileName), r)
	if err != nil {
		return "", err
	}

	err = srv.memberRepo.SetProfilePictureURL(ctx, memberID, url)
	if err != nil {
		err := srv.objectStore.Delete(ctx, path.Join("/members", fileName))
		if err != nil {
			return "", err
		}

		return "", err
	}

	return url, nil
}
