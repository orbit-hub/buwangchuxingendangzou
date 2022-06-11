package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name          string `gorm:"index:idx_mobile;unique;type:varchar(32);not null"`
	Password      string `gorm:"type:varchar(100);not null"`
	NickName      string `gorm:"type:varchar(20)"`
	FollowCount   int64  `gorm:"default:0;type:int"`
	FollowerCount int64  `gorm:"default:0;type:int"`
}
