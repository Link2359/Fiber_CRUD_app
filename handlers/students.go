package handlers

import (
	"fiber-app/database"
	"fiber-app/models"
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

	offset := (page - 1) * limit

	var total int64
	database.DB.Model(&models.Student{}).Count(&total)

	var students []models.Student
	result := database.DB.
		Order("id ASC").
		Offset(offset).
		Limit(limit).
		Find(&students)

	if result.Error != nil {
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

	if err := validators.ValidateCreateStudent(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Message,
			"field": err.Field,
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

func DeleteStudent(c *fiber.Ctx) error {
	id := c.Params("id")
	var student models.Student

	result := database.DB.First(&student, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Student not found",
		})
	}

	database.DB.Delete(&student)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Student deleted successfully",
	})
}


func UpdateStudent(c *fiber.Ctx) error {
	id := c.Params("id")
	var student models.Student

	result := database.DB.First(&student, id)
	if result.Error != nil {
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

	if req.EnrollmentNo != nil {
    	var existing models.Student
    	result := database.DB.Where("enrollment_no = ? AND id != ?", *req.EnrollmentNo, student.ID).First(&existing)
    	if result.Error == nil {
        	return c.Status(fiber.StatusConflict).JSON(fiber.Map{
            	"error": "Enrollment number already taken",
        	})
    	}
	}

	updates := map[string]any{}
	if req.Name != nil {
		updates["name"] = *req.Name
	}
	if req.EnrollmentNo != nil {
		updates["enrollment_no"] = *req.EnrollmentNo
	}
	if req.YearOfStudy != nil {
		updates["year_of_study"] = *req.YearOfStudy
	}
	if req.Department != nil {
		updates["department"] = *req.Department
	}
	if req.MobileNo != nil {
		updates["mobile_no"] = *req.MobileNo
	}

	database.DB.Model(&student).Updates(updates)

	return c.Status(fiber.StatusOK).JSON(student)
}


