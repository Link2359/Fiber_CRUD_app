package validators

import "fiber-app/models"

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

func ValidateCreateStudent(req *models.CreateStudentRequest) *ValidationError {
	if req.Name == "" {
		return &ValidationError{Field: "name", Message: "name is required"}
	}

	if len(req.Name) > 100 {
		return &ValidationError{Field: "name", Message: "name must be less than 100 characters"}
	}

	if req.EnrollmentNo == "" {
		return &ValidationError{Field: "enrollment_no", Message: "enrollment number is required"}
	}

	if len(req.EnrollmentNo) != 6 {
		return &ValidationError{Field: "enrollment_no", Message: "enrollment number must be exactly 6 characters"}
	}

	if req.Department == "" {
		return &ValidationError{Field: "department", Message: "department is required"}
	}

	if len(req.Department) > 30 {
		return &ValidationError{Field: "department", Message: "department must be less than 30 characters"}
	}

	if req.MobileNo == "" {
		return &ValidationError{Field: "mobile_no", Message: "mobile number is required"}
	}

	if len(req.MobileNo) != 10 {
		return &ValidationError{Field: "mobile_no", Message: "mobile number must be exactly 10 digits"}
	}

	if req.YearOfStudy < 1 || req.YearOfStudy > 4 {
		return &ValidationError{Field: "year_of_study", Message: "year of study must be between 1 and 4"}
	}

	return nil
}

func ValidateUpdateStudent(req *models.UpdateStudentRequest) *ValidationError {
	if req.Name != nil && len(*req.Name) > 100 {
		return &ValidationError{Field: "name", Message: "name must be less than 100 characters"}
	}

	if req.Department != nil && len(*req.Department) > 30 {
		return &ValidationError{Field: "department", Message: "department must be less than 30 characters"}
	}

	if req.EnrollmentNo != nil && len(*req.EnrollmentNo) != 6 {
		return &ValidationError{Field: "enrollment_no", Message: "enrollment number must be exactly 6 characters"}
	}

	if req.MobileNo != nil && len(*req.MobileNo) != 10 {
		return &ValidationError{Field: "mobile_no", Message: "mobile number must be exactly 10 digits"}
	}

	if req.YearOfStudy != nil && (*req.YearOfStudy < 1 || *req.YearOfStudy > 4) {
		return &ValidationError{Field: "year_of_study", Message: "year of study must be between 1 and 4"}
	}

	return nil
}

func ValidatePagination(page int, limit int) *ValidationError {
	if page < 1 {
		return &ValidationError{
			Field:   "page",
			Message: "page must be greater than 0",
		}
	}

	if limit < 1 || limit > 100 {
		return &ValidationError{
			Field:   "limit",
			Message: "limit must be between 1 and 100",
		}
	}

	if limit%5 != 0 {
		return &ValidationError{
			Field:   "limit",
			Message: "limit must be a multiple of 5",
		}
	}

	return nil
}
