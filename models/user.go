package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/ubgo/gormuuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       uuid.UUID `json:"userId" gorm:"primaryKey"`
	UserName string `json:"userName"`
	UserEmail   string `gorm:"unique" json:"userEmail"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Accounts  gormuuid.UUIDArray `gorm:"type:uuid[]"`
}
