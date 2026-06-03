package models

import "time"

type Student struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name         string    `json:"name" gorm:"not null"`
	EnrollmentNo string    `json:"enrollment_no" gorm:"unique;not null"`
	YearOfStudy  int       `json:"year_of_study" gorm:"not null"`
	Department   string    `json:"department" gorm:"not null"`
	MobileNo     string    `json:"mobile_no" gorm:"not null"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CreateStudentRequest struct {
	Name         string `json:"name" validate:"required,max=100"`
	EnrollmentNo string `json:"enrollment_no" validate:"required,len=6,numeric"`
	YearOfStudy  int    `json:"year_of_study" validate:"required,numeric,gte=1,lte=4"`
	Department   string `json:"department" validate:"required,max=50"`
	MobileNo     string `json:"mobile_no" validate:"required,len=10,numeric"`
}

type UpdateStudentRequest struct {
	Name         *string `json:"name" validate:"omitempty,max=100"`
	EnrollmentNo *string `json:"enrollment_no" validate:"omitempty,len=6,numeric"`
	YearOfStudy  *int    `json:"year_of_study" validate:"omitempty,numeric,gte=1,lte=4"`
	Department   *string `json:"department" validate:"omitempty,max=50"`
	MobileNo     *string `json:"mobile_no" validate:"omitempty,len=10,numeric"`
}
