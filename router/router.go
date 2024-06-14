package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nulhakimm/web-profile/config"
	"github.com/nulhakimm/web-profile/controller"
	"github.com/nulhakimm/web-profile/repository"
	"github.com/nulhakimm/web-profile/service"
)

func SetupRoute(app *fiber.App) {

	client, err := config.NewDatabase()
	if err != nil {
		panic(err)
	}

	experienceRepo := repository.NewExperienceRepo()
	experienceService := service.NewExperienceService(experienceRepo, client)

	skillRepo := repository.NewSkillRepo()
	skillService := service.NewSkillService(skillRepo, client)

	educationRepo := repository.NewEducationRepo()
	educationService := service.NewEducationService(educationRepo, client)

	profileRepo := repository.NewProfileRepo()
	profileService := service.NewProfileService(profileRepo, client)
	profileController := controller.NewProfileController(profileService, experienceService, skillService, educationService)

	app.Get("/", profileController.ProfileRender)

}
