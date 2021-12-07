package model

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username     string        `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Your username is required"`
	Email        string        `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required, email~Invalid email format"`
	Password     string        `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required, minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age          uint          `json:"age" form:"age" valid:"required~Your age is required, int~Age must be int, range(8|999)~Minimum age is 8"`
	Photos       []Photo       `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL;" json:"users"`
	SocialMedias []SocialMedia `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL;" json:"social_medias"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
