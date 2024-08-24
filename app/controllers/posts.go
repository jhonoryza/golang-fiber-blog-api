package controllers

import (
	"fiber_blog/app/models"
	"fiber_blog/app/responses"
	"fiber_blog/config"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func PostsIndex(c *fiber.Ctx) error {
	var posts []models.Post
	config.GetDB().Where("published_at is not null").Order("published_at desc").Find(&posts)
	postResponses := responses.NewPostResponses(&posts)
	return c.JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "OK",
		"data":    postResponses,
	})
}

func PostsShow(c *fiber.Ctx) error {
	var post models.Post
	tx := config.GetDB().Where("published_at is not null").Where("slug = ?", c.Params("slug")).First(&post)
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
