package controllers

import (
	"fiber_blog/entity"
	"fiber_blog/response"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
)

type PostController struct {
	db *gorm.DB
}

func NewPostController(db *gorm.DB) *PostController {
	return &PostController{
		db: db,
	}
}

func (controller *PostController) Index(c *fiber.Ctx) error {
	posts := []entity.Post{}
	controller.db.Where("published_at is not null").Order("published_at desc").Find(&posts)
	postResponses := response.NewPostResponses(&posts)
	return c.JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "OK",
		"data":    postResponses,
	})
}

func (controller *PostController) Show(c *fiber.Ctx) error {
	post := entity.Post{}
	tx := controller.db.Where("published_at is not null").Where("slug = ?", c.Params("slug")).First(&post)
	if tx.Error != nil {
		return c.JSON(fiber.Map{
			"code":    http.StatusNotFound,
			"message": tx.Error.Error(),
			"data":    nil,
		})
	}

	postResponse := response.NewPostResponse(&post)
	return c.JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "OK",
		"data":    postResponse,
	})
}
