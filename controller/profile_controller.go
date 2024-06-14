package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/nulhakimm/web-profile/service"
)

type ProfileController interface {
	ProfileRender(ctx *fiber.Ctx) error
}

type ProfileControllerImpl struct {
	ProfileService    service.ProfileService
	ExperienceService service.ExperienceService
}

func NewProfileController(profileService service.ProfileService, ExperienceService service.ExperienceService) ProfileController {
	return &ProfileControllerImpl{
		ProfileService:    profileService,
		ExperienceService: ExperienceService,
	}
}

func (controller *ProfileControllerImpl) ProfileRender(ctx *fiber.Ctx) error {

	profile, err := controller.ProfileService.FindProfile(ctx.Context(), "nullhakim")
	if err != nil {
		return err
	}

	experience, err := controller.ExperienceService.Find(ctx.Context())
	if err != nil {
		return fmt.Errorf("failed get all data experience : %v", err)
	}

	return ctx.Render("index", fiber.Map{
		"Profile":    profile,
		"Experience": experience,
	})

}
