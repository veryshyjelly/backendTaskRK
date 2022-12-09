package models

import "time"

type Admin struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	Name         *string   `json:"name" validate:"required"`
	Email        *string   `json:"email" validate:"email,required"`
	FacultyId    *string   `json:"faculty_id" validate:"required"`
	Contact      *string   `json:"contact"`
	Position     *string   `json:"position" validate:"required"`
	Password     *string   `json:"password" validate:"min=6"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Token        *string   `json:"token"`
	RefreshToken *string   `json:"refresh_token"`
}
