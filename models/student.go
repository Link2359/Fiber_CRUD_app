package models

import "time"

type Student struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	EnrollmentNo string    `json:"enrollment_no"`
	YearOfStudy  int       `json:"year_of_study"`
	Department   string    `json:"department"`
	MobileNo     string    `json:"mobile_no"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CreateStudentRequest struct {
	Name         string `json:"name" validate:"required,min=1,max=100"`
	EnrollmentNo string `json:"enrollment_no" validate:"required,len=6,numeric"`
	YearOfStudy  int    `json:"year_of_study" validate:"required,numeric,gte=1,lte=4"`
	Department   string `json:"department" validate:"required,min=1,max=50"`
	MobileNo     string `json:"mobile_no" validate:"required,len=10,numeric"`
}

type UpdateStudentRequest struct {
	Name         *string `json:"name" validate:"omitempty,min=1,max=100"`
	EnrollmentNo *string `json:"enrollment_no" validate:"omitempty,len=6,numeric"`
	YearOfStudy  *int    `json:"year_of_study" validate:"omitempty,numeric,gte=1,lte=4"`
	Department   *string `json:"department" validate:"omitempty,min=1,max=50"`
	MobileNo     *string `json:"mobile_no" validate:"omitempty,len=10,numeric"`
}
