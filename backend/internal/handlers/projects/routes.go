package projects

import (
    "database/sql"
    "github.com/go-redis/redis/v8"
    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(app fiber.Router, db *sql.DB, rdb *redis.Client) {
    handler := &Handler{db: db, rdb: rdb}

    app.Get("/", handler.GetAllProjects)
    app.Get("/:id", handler.GetProjectByID)
    app.Post("/", handler.CreateProject)
    app.Put("/:id", handler.UpdateProject)
    app.Delete("/:id", handler.DeleteProject)
    app.Get("/category/:category", handler.GetProjectsByCategory)
}
