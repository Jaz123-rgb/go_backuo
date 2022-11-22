package main

import (
	"context"
	"fmt"

	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/jazrgb-123/go-reactCrud/models"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	app := fiber.New()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/gomongodb"))
	if err != nil {
		panic(err)
	}

	app.Use(cors.New())

	app.Static("/", "./client/dist")

	app.Post("/users", func(c *fiber.Ctx) error {

		var user models.Users
		c.BodyParser(&user)

		//Cuando recibamos los datos
		coll := client.Database("gomongodb").Collection("users")
		result, err := coll.InsertOne(context.TODO(), bson.D{{
			Value: user.Name,
			Key:   "name",
		}})

		if err != nil {
			panic(err)
		}

		return c.JSON(&fiber.Map{
			"data": result,
		})
	})

	app.Get("/users", func(c *fiber.Ctx) error {
		var users []models.Users

		coll := client.Database("gomongodb").Collection("users")
		results, error := coll.Find(context.TODO(), bson.M{})

		if error != nil {
			panic(error)
		}

		for results.Next(context.TODO()) {
			var user models.Users
			results.Decode(&user)
			users = append(users, user)
		}
		return c.JSON(&fiber.Map{
			"users": users,
		})

	})

	app.Listen(":" + port)
	fmt.Println("Server on port 3000")
}
