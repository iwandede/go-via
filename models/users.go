package models

import (
	"time"
)

type Users struct {
	ID        int64     `gorm:"primary_key" json:"id"`
	Username  string    `gorm:"type:varchar(100);unique_index" json:"username"`
	Email     string    `gorm:"type:varchar(100);unique_index" json:"email"`
	Fullname  string    `json:"fullname"`
	Password  string    `json:"password"`
	Address   string    `json:"address"`
	Roles     string    `json:"roles"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ResponseAllUsers struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Fullname  string    `json:"fullname"`
	Address   string    `json:"address"`
	Roles     string    `json:"roles"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u Users) TableName() string {
	return "users"
}
