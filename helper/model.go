package helper

import (
	"github.com/nulhakimm/web-profile/model/domain"
	"github.com/nulhakimm/web-profile/model/web"
)

func ToProfileResponse(profile *domain.Profile) *web.ProfileResponse {

	return &web.ProfileResponse{
		Name:        profile.Name,
		Description: profile.Description,
		Email:       profile.Email,
		SocialMedia: web.SocialMedia{
			LinkedIn:  profile.SocialMedia.LinkedIn,
			Instagram: profile.SocialMedia.Instagram,
			GitHub:    profile.SocialMedia.GitHub,
		},
		Phone: profile.Phone,
		About: profile.About,
	}

}
