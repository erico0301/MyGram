package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	Message string `json:"message" form:"message" valid:"required~Message of your comment is required"`
	UserID  uint
	PhotoID uint `json:"photo_id" form:"photo_id"`
	User    *User
	Photo   *Photo
}

type GetPhotoComment struct {
	ID        uint       `json:"id"`
	Message   string     `json:"message"`
	UserID    uint       `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
}

type MyComment struct {
	ID        uint       `json:"id"`
	Message   string     `json:"message"`
	PhotoID   uint       `json:"photo_id"`
	UpdatedAt *time.Time `json:"updated_at"`
	CreatedAt *time.Time `json:"created_at"`
	Photo     struct {
		ID        uint   `json:"id"`
		Title     string `json:"titile"`
		Caption   string `json:"caption"`
		Photo_url string `json:"photo_url"`
		UserID    uint   `json:"user_id"`
	} `json:"photos"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCrete := govalidator.ValidateStruct(c)
	if errCrete != nil {
		err = errCrete
		return
	}

	err = nil
	return
}

func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(c)
	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
