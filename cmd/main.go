package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github/trace-example/server"
	"github/trace-example/storage"
	"log"
)

func main() {
	app := fiber.New()

	// Подключаемся к Redis
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	if err := client.Ping(context.TODO()).Err(); err != nil {
		log.Fatal("create redis client", err)
	}

	// Настраиваем роутер
	handler := server.NewFiberHandler(storage.NewNotesStorage(client))
	app.Post("/create", handler.CreateNote)
	app.Get("/get", handler.GetNote)

	log.Fatal(app.Listen(":8080"))
}
