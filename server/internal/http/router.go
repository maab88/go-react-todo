package http

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"todoapp/internal/models"
)

type Server struct {
	app *fiber.App
	db  *gorm.DB
}

func NewServer(sqlitePath string) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok { code = e.Code }
			return c.Status(code).JSON(fiber.Map{"error": err.Error()})
		},
	})

	app.Use(cors.New())

	dsn := sqlitePath
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := db.AutoMigrate(&models.Task{}); err != nil {
		log.Fatalf("auto-migrate failed: %v", err)
	}

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"ok": true, "time": time.Now()})
	})

	api := app.Group("/api")
	registerTaskRoutes(api, db)

	return app
}
