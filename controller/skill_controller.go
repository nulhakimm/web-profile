package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nulhakimm/web-profile/service"
)

type SkillController interface {
	Save(ctx *fiber.Ctx) error
}

type SkillControllerImpl struct {
	SkillService service.SkillService
}

func NewSkillController(skillService service.SkillService) SkillController {
	return &SkillControllerImpl{
		SkillService: skillService,
	}
}

func (controller *SkillControllerImpl) Save(ctx *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}
