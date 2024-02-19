package controller

import (
	"backend-project/database"
	"backend-project/models"
	"backend-project/util"
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreatePost(c *fiber.Ctx) error {
	var blogpost models.Blog
	if err := c.BodyParser(&blogpost); err != nil {
		fmt.Println("Error parsing request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Payload",
			"error":   err.Error(),
		})
	}

	userID := c.Locals("userID").(string)

	blogpost.UserID = userID

	// Create post in the database
	if err := database.DB.Create(&blogpost).Error; err != nil {
		fmt.Println("Error creating post in the database:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to create post",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Success post",
	})
}


func AllPost(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid page number",
		})
	}
	limit := 5
	offset := (page - 1) * limit
	var total int64
	var getblog []models.Blog
	database.DB.Preload("User").Offset(offset).Limit(limit).Find(&getblog)
	database.DB.Model(&models.Blog{}).Count(&total)
	return c.JSON(fiber.Map{
		"data" : getblog,
		"meta" : fiber.Map{
			"total" : total,
			"page" : page,
			"last_page":math.Ceil(float64(total)/float64(limit)),
		},
	})
}

func DetailPost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var blogpost models.Blog
	database.DB.Where("id=?", id).Preload("User").First(&blogpost)
	return c.JSON((fiber.Map{
		"data": blogpost,
	}))
}

func UpdatePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	blog := models.Blog{
		Id:uint(id),
	}
	if err := c.BodyParser(&blog) ; err != nil {
		fmt.Println("Unable to parse body", err)
	}
	database.DB.Model(&blog).Updates(blog)
	return c.JSON(fiber.Map{
		"message": "Success update post",
		"data": blog,
	})
}

func UniquePost(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	id, _ := util.ParseJwt(cookie)
	var blog []models.Blog
	database.DB.Model(&blog).Where("user_id = ?", id).Preload("User").Find(&blog)
	return c.JSON(fiber.Map{
		"data": blog,
	})
}

func DeletePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	blog := models.Blog{
		Id: uint(id),
	}
	deleteQuery := database.DB.Delete(&blog)
	if errors.Is(deleteQuery.Error, gorm.ErrRecordNotFound) {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message" : "Oppss! Blog not found",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Success delete post",
	})
}