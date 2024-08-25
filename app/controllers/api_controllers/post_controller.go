package api_controllers

import (
	"fiber_blog/app/models"
	"fiber_blog/app/responses"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
)

func PostIndex(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var posts []models.Post
		db.Where("published_at is not null").Order("published_at desc").Find(&posts)
		postResponses := responses.NewPostResponses(&posts)
		return c.JSON(fiber.Map{
			"code":    http.StatusOK,
			"message": "OK",
			"data":    postResponses,
		})
	}
}

func PostShow(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var post models.Post
		tx := db.Where("published_at is not null").Where("slug = ?", c.Params("slug")).First(&post)
		if tx.Error != nil {
			return c.JSON(fiber.Map{
				"code":    http.StatusNotFound,
				"message": tx.Error.Error(),
				"data":    nil,
			})
		}

		postResponse := responses.NewPostResponse(&post)
		return c.JSON(fiber.Map{
			"code":    http.StatusOK,
			"message": "OK",
			"data":    postResponse,
		})
	}
}
