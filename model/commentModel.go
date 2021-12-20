package model

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	GormModel
	Message   string    `json:"message" form:"message" valid:"required~Message of your comment is required"`
	UserID    uint      `json:"user_id"`
	PhotoID   uint      `json:"photo_id"`
	User      *User     `gorm:"foreignKey:UserID"`
	Photo     *Photo    `gorm:"foreignKey:PhotoID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
