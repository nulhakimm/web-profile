package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nulhakimm/web-profile/service"
)

type EducationController interface {
	Create(ctx *fiber.Ctx) error
}

type EducationControllerImpl struct {
	EducationService service.EducationService
}

func NewEducationController(educationService service.EducationService) EducationController {
	return &EducationControllerImpl{
		EducationService: educationService,
	}
}

func (controller *EducationControllerImpl) Create(ctx *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}
