package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	handler "github.com/sushant102004/aqua-watch-backend/internal/app/handlers"
	"github.com/sushant102004/aqua-watch-backend/internal/app/store"
	db "github.com/sushant102004/aqua-watch-backend/internal/database"
)

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	app := fiber.New(config)
	client := db.ConnectToMongo("mongodb+srv://sushantdhiman9812:sushant@aquawatchcluster.h6h1uow.mongodb.net/?retryWrites=true&w=majority")

	userHandler := handler.NewUserHandler(store.NewMongoUserStore(client))
	postHandler := handler.NewPostHandler(store.NewMongoPostStore(client))
	ngoHandler := handler.NewNGOHandler(store.NewMongoNGOStore(client))

	app.Post("/user", userHandler.HandleCreateUser)
	app.Get("/login", userHandler.HandleLoginUSer)

	app.Post("/post", postHandler.HandleInsertPost)
	app.Get("/post", postHandler.HandleGetAllPosts)
	app.Get("/post/search", postHandler.HandleSearchPostsViaCity)

	app.Post("/post/damage-score", postHandler.HandleUpdateDamageScore)
	app.Get("/post/map", postHandler.HandleGetPostsForMap)

	app.Post("/ngo/signup", ngoHandler.HandleSignUp)
	app.Get("/ngo/login", ngoHandler.HandleLogin)

	log.Fatal(app.Listen(":5000"))
}
