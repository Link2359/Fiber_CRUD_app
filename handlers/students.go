package handlers

import (
	"fiber-app/database"
	"fiber-app/models"

	"github.com/gofiber/fiber/v2"
)


func GetAllStudents(c *fiber.Ctx) error {
	var students []models.Student

	result := database.DB.Find(&students)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch students",
		})
	}

	return c.JSON(fiber.Map{
		"data":  students,
		"count": len(students),
	})
}


func GetStudent(c *fiber.Ctx) error {
	id := c.Params("id")
	var student models.Student

	result := database.DB.First(&student, id)
	if result.Error != nil {
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

	if req.Name == "" || req.EnrollmentNo == "" || req.Department == "" || req.MobileNo == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Name, EnrollmentNo, Department and MobileNo are required fields",
		})
	}


	if len(req.EnrollmentNo) != 6 {
    		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
        		"error": "Enrollment number must be exactly 6 characters",
    		})
	}

	if len(req.MobileNo) != 10 {
    		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
        		"error": "Mobile number must be exactly 10 characters",
    		})
	}


	student := models.Student{
		Name:         req.Name,
		EnrollmentNo: req.EnrollmentNo,
		YearOfStudy:  req.YearOfStudy,
		Department:   req.Department,
		MobileNo:     req.MobileNo,
	}

	result := database.DB.Create(&student)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create student",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(student)
}