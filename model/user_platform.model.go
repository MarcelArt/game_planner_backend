package model

import "gorm.io/gorm"

const userPlatformTableName = "user_platforms"

type UserPlatform struct {
	gorm.Model
	Type      string `gorm:"not null" json:"type"`
	Username  string `gorm:"not null" json:"username"`
	ProfileID uint   `gorm:"not null" json:"profileId"`

	Profile *Profile `json:"-"`
}

type UserPlatformInput struct {
	ModelInput
	Type      string `gorm:"not null" json:"type"`
	Username  string `gorm:"not null" json:"username"`
	ProfileID uint   `gorm:"not null" json:"profileId"`

	Profile *Profile `json:"-"`
}

type UserPlatformResponse struct {
	gorm.Model
	Type      string `gorm:"not null" json:"type"`
	Username  string `gorm:"not null" json:"username"`
	ProfileID uint   `gorm:"not null" json:"profileId"`

	Profile *Profile `json:"profile"`
}

func (UserPlatformInput) TableName() string {
	return userPlatformTableName
}

func (UserPlatformResponse) TableName() string {
	return userPlatformTableName
}
