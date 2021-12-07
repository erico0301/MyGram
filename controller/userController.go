package controller

import (
	"MyGram/database"
	"MyGram/helper"
	"MyGram/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)
	User := model.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         User.ID,
		"age":        User.Age,
		"email":      User.Email,
		"password":   User.Password,
		"username":   User.Username,
		"created_at": User.CreatedAt,
	})
}

func UserLogin(c *gin.Context) {

}
