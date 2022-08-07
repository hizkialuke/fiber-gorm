package models

import "time"

type User struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	UserName  string    `json:"user_name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy int64     `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int64     `json:"updated_by"`
}

type LoginRequest struct {
	UserName string `json:"user_name" validate:"required"`
	Password string `json:"password" validate:"required"`
}
