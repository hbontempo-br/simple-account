package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type OperationType struct {
	gorm.Model
	ID         uint `gorm:"primary_key"`
	Enumerator string `gorm:"unique;not null"`
	Description string `gorm:"unique;not null"`
	CreateAt     *time.Time

}

func (OperationType) TableName() string {
	return "Account"
}