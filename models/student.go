package models

import "time"

type Student struct {
	ID             uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name           string    `json:"name" gorm:"not null"`
	EnrollmentNo   string    `json:"enrollment_no" gorm:"unique;not null"`
	YearOfStudy    int       `json:"year_of_study" gorm:"not null"`
	Department     string    `json:"department" gorm:"not null"`
	MobileNo       string    `json:"mobile_no" gorm:"not null"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type CreateStudentRequest struct {
	Name         string `json:"name"`
	EnrollmentNo string `json:"enrollment_no"`
	YearOfStudy  int    `json:"year_of_study"`
	Department   string `json:"department"`
	MobileNo     string `json:"mobile_no"`
}

type UpdateStudentRequest struct {
	Name         *string `json:"name"`
	EnrollmentNo *string `json:"enrollment_no"`
	YearOfStudy  *int    `json:"year_of_study"`
	Department   *string `json:"department"`
	MobileNo     *string `json:"mobile_no"`
}