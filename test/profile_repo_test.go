package test

import (
	"context"
	"testing"

	"github.com/nulhakimm/web-profile/config"
	"github.com/nulhakimm/web-profile/model/web"
	"github.com/nulhakimm/web-profile/repository"
	"github.com/nulhakimm/web-profile/service"
	"github.com/stretchr/testify/assert"
)

var (
	profileRepo    = repository.NewProfileRepo()
	client, _      = config.NewDatabase()
	profileService = service.NewProfileService(profileRepo, client)
)

func TestCreateProfile(t *testing.T) {

	profile := &web.ProfileCreate{
		Name:        "nullhakim",
		Description: "hello i'am fresh graduate",
		Email:       "nullhakim@mail.com",
		SocialMedia: web.SocialMedia{
			LinkedIn:  "linkedin.com/nullhakim",
			Instagram: "instagram.com/nullhakim",
			GitHub:    "github.com/nullhakim",
		},
		Phone: "289-2594-4458-554",
		About: `Lorem Ipsum is simply dummy text of the 
		printing and typesetting industry. 
		Lorem Ipsum has been the industry's standard
		dummy text ever since the 1500s,
		when an unknown printer took a galley of type and 
		scrambled it to make a type specimen book`,
	}

	err := profileService.Save(context.Background(), profile)
	if err != nil {
		panic(err)
	}
	assert.Nil(t, err)

}
