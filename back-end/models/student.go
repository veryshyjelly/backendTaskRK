package models

import "time"

type Student struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	Name         *string   `json:"name" validate:"required"`
	RollNo       *string   `json:"roll_no" validate:"required"`
	Email        *string   `json:"email" validate:"email,required"`
	Password     *string   `json:"password" validate:"min=6,required"`
	Contact      *string   `json:"contact"`
	BlockNo      *string   `json:"block_no"`
	RoomNo       *string   `json:"room_no"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Token        *string   `json:"token"`
	RefreshToken *string   `json:"refresh_token"`
}
