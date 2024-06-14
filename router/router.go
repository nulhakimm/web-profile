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
	serviceRepo := service.NewExperienceService(experienceRepo, client)

	profileRepo := repository.NewProfileRepo()
	profileService := service.NewProfileService(profileRepo, client)
	profileController := controller.NewProfileController(profileService, serviceRepo)

	app.Get("/", profileController.ProfileRender)

}
