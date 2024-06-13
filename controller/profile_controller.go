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
	SkillService   service.SkillService
}

func NewProfileController(profileService service.ProfileService, skillService service.SkillService) ProfileController {
	return &ProfileControllerImpl{
		ProfileService: profileService,
		SkillService:   skillService,
	}
}

func (controller *ProfileControllerImpl) ProfileRender(ctx *fiber.Ctx) error {

	profile, err := controller.ProfileService.FindProfile(ctx.Context(), "nullhakim")
	if err != nil {
		return err
	}

	skill, err := controller.SkillService.Find(ctx.Context())
	if err != nil {
		return fmt.Errorf("failed get all data skill : %v", err)
	}

	return ctx.Render("index", fiber.Map{
		"Profile": profile,
		"Skill":   skill,
	})

}
