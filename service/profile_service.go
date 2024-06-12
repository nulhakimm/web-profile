package service

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/nulhakimm/web-profile/helper"
	"github.com/nulhakimm/web-profile/model/domain"
	"github.com/nulhakimm/web-profile/model/web"
	"github.com/nulhakimm/web-profile/repository"
)

type ProfileService interface {
	FindProfile(ctx context.Context, name string) (*web.ProfileResponse, error)
	Save(ctx context.Context, request *web.ProfileCreate) error
}

type ProfileServiceImpl struct {
	ProfileRepo repository.ProfileRepo
	Client      *firestore.Client
}

func NewProfileService(profileRepo repository.ProfileRepo, client *firestore.Client) ProfileService {
	return &ProfileServiceImpl{
		ProfileRepo: profileRepo,
		Client:      client,
	}
}

func (service *ProfileServiceImpl) FindProfile(ctx context.Context, name string) (*web.ProfileResponse, error) {

	profileResponse, err := service.ProfileRepo.FindProfile(ctx, service.Client, name)
	if err != nil {
		return nil, err
	}

	return helper.ToProfileResponse(profileResponse), nil

}

func (service *ProfileServiceImpl) Save(ctx context.Context, request *web.ProfileCreate) error {

	profile := &domain.Profile{
		Name:        request.Name,
		Description: request.Description,
		Email:       request.Email,
		SocialMedia: domain.SocialMedia{
			LinkedIn:  request.SocialMedia.LinkedIn,
			Instagram: request.SocialMedia.Instagram,
			GitHub:    request.SocialMedia.GitHub,
		},
		Phone: request.Phone,
		About: request.About,
	}

	err := service.ProfileRepo.Save(ctx, service.Client, profile)
	if err != nil {
		return err
	}

	return nil

}
