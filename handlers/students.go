package handlers

import (
	"fiber-app/models"
	"fiber-app/repository"
	"fiber-app/validators"

	"github.com/gofiber/fiber/v2"
)

func GetAllStudents(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 5)

	if err := validators.ValidatePagination(page, limit); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Message,
			"field": err.Field,
		})
	}

	students, total, err := repository.GetAllStudents(page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch students",
		})
	}

	totalPages := (int(total) + limit - 1) / limit

	return c.JSON(fiber.Map{
		"data":        students,
		"total":       total,
		"page":        page,
		"limit":       limit,
		"total_pages": totalPages,
	})
}

func GetStudent(c *fiber.Ctx) error {
	id := c.Params("id")

	student, err := repository.GetStudentByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Student not found",
		})
	}

	return c.JSON(student)
}

func CreateStudent(c *fiber.Ctx) error {
	req := new(models.CreateStudentRequest)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := validators.ValidateCreateStudent(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Message,
			"field": err.Field,
		})
	}

	student, err := repository.CreateStudent(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create student",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(student)
}

func DeleteStudent(c *fiber.Ctx) error {
	id := c.Params("id")

	student, err := repository.GetStudentByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Student not found",
		})
	}

	if err := repository.DeleteStudent(student); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete student",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Student deleted successfully",
	})
}

func UpdateStudent(c *fiber.Ctx) error {
	id := c.Params("id")

	student, err := repository.GetStudentByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Student not found",
		})
	}

	req := new(models.UpdateStudentRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := validators.ValidateUpdateStudent(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Message,
			"field": err.Field,
		})
	}

	if req.EnrollmentNo != nil && repository.IsEnrollmentNoTaken(*req.EnrollmentNo, student.ID) {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Enrollment number already taken",
		})
	}

	updated, err := repository.UpdateStudent(student, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update student",
		})
	}

	return c.Status(fiber.StatusOK).JSON(updated)
}
