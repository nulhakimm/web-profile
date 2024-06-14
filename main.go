package main

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/nulhakimm/web-profile/config"
	"github.com/nulhakimm/web-profile/model/web"
	"github.com/nulhakimm/web-profile/repository"
	"github.com/nulhakimm/web-profile/router"
	"github.com/nulhakimm/web-profile/service"
)

func Create() {

	db, err := config.NewDatabase()
	if err != nil {
		panic(err)
	}

	repo := repository.NewEducationRepo()
	service := service.NewEducationService(repo, db)

	model := &web.EducationCreate{
		Name:      "OXFORD Univ",
		Address:   "Jl. Address 1",
		EntryYear: "2019",
		OutYear:   "2023",
		Title:     "Bahcelor Deegre of informatic",
		Achiev: []string{
			"graduate on time",
			"make some app",
			"graduate for perfect score",
		},
	}

	service.Save(context.Background(), model)

}

func main() {

	// Create()

	engine := html.New("./public/views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	router.SetupRoute(app)

	app.Listen("localhost:9090")

}
