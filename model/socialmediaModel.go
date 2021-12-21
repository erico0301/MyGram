package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name           string `json:"name" form:"name" valid:"required~Name of your social media is required"`
	SocialMediaURL string `json:"social_media_url" form:"social_medai_url" valid:"required~ Socialmedai URL of yout Socialmedia is required"`
	UserID         uint
	User           *User
}

type MySocialMedia struct {
	ID             uint       `json:"id"`
	Name           string     `json:"name"`
	SocialMediaURL string     `json:"social_media_url"`
	UserID         uint       `json:"user_id"`
	CreatedAt      *time.Time `json:"created_at"`
	User           struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	} `json:"user"`
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
func (s *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
