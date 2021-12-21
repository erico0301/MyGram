package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title     string `json:"title" form:"title" valid:"required~Title of your photo is required"`
	Caption   string `json:"caption" form:"caption"`
	Photo_url string `json:"photo_url" form:"photo_url" valid:"required~Photo URL of your photo is required"`
	UserID    uint
	User      *User
	Comments  []Comment `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL;" json:"comments"`
}

type MyPhoto struct {
	ID        uint              `json:"id"`
	Title     string            `json:"title"`
	Caption   string            `json:"caption"`
	Photo_url string            `json:"photo_url"`
	CreatedAt *time.Time        `json:"created_at"`
	Comments  []GetPhotoComment `json:"comments"`
}
type PublicPhoto struct {
	ID        uint       `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	Photo_url string     `json:"photo_url"`
	CreatedAt *time.Time `json:"created_at"`
	User      struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}
	Comments []GetPhotoComment `json:"comments"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(p)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
