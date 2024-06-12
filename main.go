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

	// repo := repository.NewProfileRepo()
	// service := service.NewProfileService(repo, db)

	// profile := &web.ProfileCreate{
	// 	Name:        "nullhakim",
	// 	Description: "hello i'am fresh graduate",
	// 	Email:       "nullhakim@mail.com",
	// 	SocialMedia: web.SocialMedia{
	// 		LinkedIn:  "linkedin.com/nullhakim",
	// 		Instagram: "instagram.com/nullhakim",
	// 		GitHub:    "github.com/nullhakim",
	// 	},
	// 	Phone: "289-2594-4458-554",
	// 	About: `Lorem Ipsum is simply dummy text of the
	// 	printing and typesetting industry.
	// 	Lorem Ipsum has been the industry's standard
	// 	dummy text ever since the 1500s,
	// 	when an unknown printer took a galley of type and
	// 	scrambled it to make a type specimen book`,
	// }

	// err = service.Save(context.Background(), profile)
	// if err != nil {
	// 	panic(err)
	// }

	engine := html.New("./public/views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	router.SetupRoute(app)

	app.Listen("localhost:9090")

}
