package model

import "gorm.io/gorm"

const profileTableName = "profile"

type Profile struct {
	gorm.Model
	Name   string `gorm:"not null" json:"name"`
	Avatar string `json:"avatar"`
	UserID uint   `gorm:"not null" json:"userId"`

	User *User `json:"-"`
}

type ProfileInput struct {
	ModelInput
	Name   string `gorm:"not null" json:"name"`
	Avatar string `json:"avatar"`
	UserID uint   `gorm:"not null" json:"userId"`

	User *User `json:"-"`
}

type ProfileResponse struct {
	gorm.Model
	Name   string `gorm:"not null" json:"name"`
	Avatar string `json:"avatar"`
	UserID uint   `gorm:"not null" json:"userId"`

	User *User `json:"user"`
}

func (ProfileInput) TableName() string {
	return profileTableName
}

func (ProfileResponse) TableName() string {
	return profileTableName
}
