package models

import (
	"time"
	"gorm.io/gorm"
	"github.com/ubgo/gormuuid"
)

type User struct {
	gorm.Model
	Id       string `json:"userId"`
	UserName string `json:"userName"`
	UserEmail   string `gorm:"unique" json:"userEmail"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Accounts  gormuuid.UUIDArray `gorm:"type:uuid[]"`
}
