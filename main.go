package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/nulhakimm/web-profile/router"
)

func main() {

	// db, err := config.NewDatabase()
	// if err != nil {
	// 	panic(err)
	// }

	// skillRepo := repository.NewSkillRepo()
	// skillService := service.NewSkillService(skillRepo, db)

	// skill := &web.SkillCreate{
	// 	Name: "postgreSQL",
	// }

	// skillService.Save(context.Background(), skill)

	engine := html.New("./public/views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	router.SetupRoute(app)

	app.Listen("localhost:9090")

}
