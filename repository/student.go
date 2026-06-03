package repository

import (
	"fiber-app/database"
	"fiber-app/models"
)

func GetAllStudents(page int, limit int) ([]models.Student, int64, error) {
	var students []models.Student
	var total int64

	database.DB.Model(&models.Student{}).Count(&total)

	offset := (page - 1) * limit

	result := database.DB.
		Order("id ASC").
		Offset(offset).
		Limit(limit).
		Find(&students)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	return students, total, nil
}

func GetStudentByID(id string) (*models.Student, error) {
	var student models.Student

	result := database.DB.First(&student, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &student, nil
}

func CreateStudent(req *models.CreateStudentRequest) (*models.Student, error) {
	student := models.Student{
		Name:         req.Name,
		EnrollmentNo: req.EnrollmentNo,
		YearOfStudy:  req.YearOfStudy,
		Department:   req.Department,
		MobileNo:     req.MobileNo,
	}

	result := database.DB.Create(&student)
	if result.Error != nil {
		return nil, result.Error
	}

	return &student, nil
}

func IsEnrollmentNoTaken(enrollmentNo string, excludeID uint) bool {
	var existing models.Student
	result := database.DB.
		Where("enrollment_no = ? AND id != ?", enrollmentNo, excludeID).
		First(&existing)
	return result.Error == nil
}

func UpdateStudent(student *models.Student, req *models.UpdateStudentRequest) (*models.Student, error) {
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

	result := database.DB.Model(student).Updates(updates)
	if result.Error != nil {
		return nil, result.Error
	}

	return student, nil
}

func DeleteStudent(student *models.Student) error {
	result := database.DB.Delete(student)
	return result.Error
}
