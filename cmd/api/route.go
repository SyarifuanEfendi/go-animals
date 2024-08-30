package api

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/SyarifuanEfendi/go-animals/internal/handlers"
	"github.com/SyarifuanEfendi/go-animals/internal/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Route() {
    app := fiber.New()

    app.Use(recover.New())
    app.Use(logger.New())

    db, err := storage.NewPostgresDB()
    if err != nil {
        log.Fatalf("Could not connect to database: %v", err)
    }
    defer db.Close()

	if err := createTables(db); err != nil {
        log.Fatalf("Could not create tables: %v", err)
    }

    app.Get("/animals", handlers.GetAnimals(db))
    app.Get("/animals/:id", handlers.GetAnimalByID(db))
    app.Post("/animals", handlers.CreateAnimal(db))
    app.Put("/animals/:id", handlers.UpdateAnimal(db))
    app.Delete("/animals/:id", handlers.DeleteAnimal(db))

    log.Fatal(app.Listen(":8080"))
}

func createTables(db *sql.DB) error {
    _, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS animals (
            id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            class VARCHAR(255) NOT NULL,
            legs INT NOT NULL
        );
    `)
    if err != nil {
        return fmt.Errorf("could not create tables: %w", err)
    }
    return nil
}