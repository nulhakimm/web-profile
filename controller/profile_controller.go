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
	ProfileService service.ProfileService
}

func NewProfileController(profileService service.ProfileService) ProfileController {
	return &ProfileControllerImpl{
		ProfileService: profileService,
	}
}

func (controller *ProfileControllerImpl) ProfileRender(ctx *fiber.Ctx) error {

	profile, err := controller.ProfileService.FindProfile(ctx.Context(), "nullhakim")
	if err != nil {
		return fmt.Errorf("failed get all data : %v", err)
	}

	return ctx.Render("index", profile)

}
