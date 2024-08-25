package web_controllers

import (
	"fiber_blog/app/models"
	"fiber_blog/app/responses"
	"github.com/gofiber/fiber/v2"
	"github.com/jhonoryza/inertia-fiber"
	"gorm.io/gorm"
)

func ArticleIndex(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var posts []models.Post
		db.Where("published_at is not null").Order("published_at desc").Find(&posts)
		postResponses := responses.NewPostResponses(&posts)

		return inertia.Render(c, 200, "Post/Index", fiber.Map{
			"posts": postResponses,
		})
	}
}

func ArticleShow(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var post models.Post
		tx := db.Where("published_at is not null").Where("slug = ?", c.Params("slug")).First(&post)
		if tx.Error != nil {
			return inertia.Render(c, 404, "Error", fiber.Map{
				"message": "Article Not Found",
			})
		}

		postResponse := responses.NewPostResponse(&post)

		return inertia.Render(c, 200, "Post/Show", fiber.Map{
			"post": postResponse,
		})
	}
}
