package handlers

import (
	"database/sql"

	"github.com/SyarifuanEfendi/go-animals/internal/dto"
	"github.com/SyarifuanEfendi/go-animals/internal/helper"
	"github.com/SyarifuanEfendi/go-animals/internal/models"
	"github.com/SyarifuanEfendi/go-animals/internal/storage"

	"github.com/gofiber/fiber/v2"
)

func GetAnimals(db *sql.DB) fiber.Handler {
    return func(c *fiber.Ctx) error {
        rows, err := db.Query("SELECT id, name, class, legs FROM animals")
        if err != nil {
            return fiber.NewError(fiber.StatusInternalServerError, err.Error())
        }
        defer rows.Close()

        var animals []dto.AnimalResponse
        for rows.Next() {
            var a dto.AnimalResponse
            if err := rows.Scan(&a.ID, &a.Name, &a.Class, &a.Legs); err != nil {
                return fiber.NewError(fiber.StatusInternalServerError, err.Error())
            }
            animals = append(animals, a)
        }
        if len(animals) == 0 {
            return fiber.NewError(fiber.StatusNotFound, "No animals found")
        }
        return c.JSON(animals)
    }
}

func GetAnimalByID(db *sql.DB) fiber.Handler {
    return func(c *fiber.Ctx) error {
        id, err := helper.ParseIntFromParam(c, "id")
        if err != nil {
            return err
        }

        row := db.QueryRow("SELECT id, name, class, legs FROM animals WHERE id = $1", id)
        var a dto.AnimalResponse
        if err := row.Scan(&a.ID, &a.Name, &a.Class, &a.Legs); err != nil {
            if err == sql.ErrNoRows {
                return fiber.NewError(fiber.StatusNotFound, "Animal not found")
            }
            return fiber.NewError(fiber.StatusInternalServerError, err.Error())
        }
        return c.JSON(a)
    }
}

func CreateAnimal(db *sql.DB) fiber.Handler {
    return func(c *fiber.Ctx) error {
        var a models.Animal
        if err := c.BodyParser(&a); err != nil {
            return fiber.NewError(fiber.StatusBadRequest, err.Error())
        }

        tx, err := storage.BeginTransaction(db)
        if err != nil {
            return fiber.NewError(fiber.StatusInternalServerError, err.Error())
        }

        var exists bool
        err = tx.QueryRow("SELECT EXISTS(SELECT 1 FROM animals WHERE name = $1)", a.Name).Scan(&exists)
        if err != nil {
            storage.RollbackTransaction(tx)
            return fiber.NewError(fiber.StatusInternalServerError, err.Error())
        }
        if exists {
            storage.RollbackTransaction(tx)
            return fiber.NewError(fiber.StatusConflict, "Animal with this Name already exists")
        }

        _, err = tx.Exec("INSERT INTO animals (name, class, legs) VALUES ($1, $2, $3)", a.Name, a.Class, a.Legs)
        if err != nil {
            storage.RollbackTransaction(tx)
            return fiber.NewError(fiber.StatusInternalServerError, err.Error())
        }

        if err := storage.CommitTransaction(tx); err != nil {
            return fiber.NewError(fiber.StatusInternalServerError, err.Error())
        }

        return c.SendStatus(fiber.StatusCreated)
    }
}

func UpdateAnimal(db *sql.DB) fiber.Handler {
    return func(c *fiber.Ctx) error {
        id, err := helper.ParseIntFromParam(c, "id")
        if err != nil {
            return err
        }

        var a models.Animal
        if err := c.BodyParser(&a); err != nil {
            return fiber.NewError(fiber.StatusBadRequest, err.Error())
        }

        tx, err := storage.BeginTransaction(db)
        if err != nil {
            return fiber.NewError(fiber.StatusInternalServerError, err.Error())
        }

        var exists bool
        err = tx.QueryRow("SELECT EXISTS(SELECT 1 FROM animals WHERE id != $1 AND name = $2)", id, a.Name).Scan(&exists)
        if err != nil {
            storage.RollbackTransaction(tx)
            return fiber.NewError(fiber.StatusInternalServerError, err.Error())
        }
        if exists {
            storage.RollbackTransaction(tx)
            return fiber.NewError(fiber.StatusConflict, "Animal with this Name already exists")
        }

        _, err = tx.Exec("UPDATE animals SET name = $1, class = $2, legs = $3 where id = $4", a.Name, a.Class, a.Legs, id)
        if err != nil {
            storage.RollbackTransaction(tx)
            return fiber.NewError(fiber.StatusInternalServerError, err.Error())
        }

        if err := storage.CommitTransaction(tx); err != nil {
            return fiber.NewError(fiber.StatusInternalServerError, err.Error())
        }

        return c.SendStatus(fiber.StatusOK)
    }
}

func DeleteAnimal(db *sql.DB) fiber.Handler {
    return func(c *fiber.Ctx) error {
        id, err := helper.ParseIntFromParam(c, "id")
        if err != nil {
            return err
        }

        tx, err := storage.BeginTransaction(db)
        if err != nil {
            return fiber.NewError(fiber.StatusInternalServerError, err.Error())
        }

        res, err := tx.Exec("DELETE FROM animals WHERE id = $1", id)
        if err != nil {
            storage.RollbackTransaction(tx)
            return fiber.NewError(fiber.StatusInternalServerError, err.Error())
        }

        rowsAffected, err := res.RowsAffected()
        if err != nil {
            storage.RollbackTransaction(tx)
            return fiber.NewError(fiber.StatusInternalServerError, err.Error())
        }
        if rowsAffected == 0 {
            storage.RollbackTransaction(tx)
            return fiber.NewError(fiber.StatusNotFound, "Animal not found")
        }

        if err := storage.CommitTransaction(tx); err != nil {
            return fiber.NewError(fiber.StatusInternalServerError, err.Error())
        }

        return c.SendStatus(fiber.StatusOK)
    }
}