package http

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"todoapp/internal/models"
)

type createTaskDTO struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	DueDate     *time.Time `json:"dueDate"`
	Priority    int        `json:"priority"` // 1=high,2=med,3=low
}

type updateTaskDTO struct {
	Title       *string     `json:"title"`
	Description *string     `json:"description"`
	DueDate     **time.Time `json:"dueDate"`
	Priority    *int        `json:"priority"`
	Status      *string     `json:"status"` // todo|doing|done
}

func registerTaskRoutes(r fiber.Router, db *gorm.DB) {
	tasks := r.Group("/tasks")

	tasks.Get("/", func(c *fiber.Ctx) error {
		var list []models.Task
		status := c.Query("status")
		q := c.Query("q")

		query := db.Order("created_at DESC")
		if status != "" {
			query = query.Where("status = ?", status)
		}
		if q != "" {
			like := "%%" + q + "%%"
			query = query.Where("title LIKE ? OR description LIKE ?", like, like)
		}
		if err := query.Find(&list).Error; err != nil {
			return err
		}
		return c.JSON(fiber.Map{"data": list})
	})

	tasks.Get(":id", func(c *fiber.Ctx) error {
		var t models.Task
		if err := db.First(&t, "id = ?", c.Params("id")).Error; err != nil {
			return fiber.NewError(fiber.StatusNotFound, "task not found")
		}
		return c.JSON(fiber.Map{"data": t})
	})

	tasks.Post("/", func(c *fiber.Ctx) error {
		var in createTaskDTO
		if err := c.BodyParser(&in); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		if in.Title == "" {
			return fiber.NewError(fiber.StatusBadRequest, "title is required")
		}
		if in.Priority == 0 { in.Priority = 2 }

		t := models.Task{
			Title: in.Title,
			Description: in.Description,
			DueDate: in.DueDate,
			Priority: in.Priority,
			Status: "todo",
		}
		if err := db.Create(&t).Error; err != nil {
			return err
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"data": t})
	})

	tasks.Put(":id", func(c *fiber.Ctx) error {
		var in updateTaskDTO
		if err := c.BodyParser(&in); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		var t models.Task
		if err := db.First(&t, "id = ?", c.Params("id")).Error; err != nil {
			return fiber.NewError(fiber.StatusNotFound, "task not found")
		}
		if in.Title != nil { t.Title = *in.Title }
		if in.Description != nil { t.Description = *in.Description }
		if in.DueDate != nil { t.DueDate = *in.DueDate }
		if in.Priority != nil { t.Priority = *in.Priority }
		if in.Status != nil { t.Status = *in.Status }
		if err := db.Save(&t).Error; err != nil {
			return err
		}
		return c.JSON(fiber.Map{"data": t})
	})

	tasks.Patch(":id/status", func(c *fiber.Ctx) error {
		var in struct{ Status string `json:"status"` }
		if err := c.BodyParser(&in); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		if in.Status == "" { return fiber.NewError(fiber.StatusBadRequest, "status is required") }

		var t models.Task
		if err := db.First(&t, "id = ?", c.Params("id")).Error; err != nil {
			return fiber.NewError(fiber.StatusNotFound, "task not found")
		}
		t.Status = in.Status
		if err := db.Save(&t).Error; err != nil {
			return err
		}
		return c.JSON(fiber.Map{"data": t})
	})

	tasks.Delete(":id", func(c *fiber.Ctx) error {
		if err := db.Delete(&models.Task{}, "id = ?", c.Params("id")).Error; err != nil {
			return err
		}
		return c.SendStatus(fiber.StatusNoContent)
	})
}
