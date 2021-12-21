package controller

import (
	"MyGram/database"
	"MyGram/helper"
	"MyGram/model"
	"fmt"
	"net/http"
	"strconv"

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
	db := database.GetDB()
	contentType := helper.GetContentType(c)
	_, _ = db, contentType

	User := model.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password
	fmt.Println("before:\n", User)
	err := db.Debug().Where("email=?", User.Email).Take(&User).Error
	fmt.Println("after:\n", User)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	comparePass := helper.ComparePass([]byte(User.Password), []byte(password))
	fmt.Println("Compare Password : ", comparePass)
	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	token := helper.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}

func UserEdit(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)
	User := model.User{}
	userID, _ := strconv.Atoi(c.Param("user_id"))
	myUser := model.User{}
	_ = db.First(&myUser, userID)

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	User.Username = myUser.Username
	User.Age = myUser.Age
	User.Password = helper.HashPass(User.Password)

	err := db.Model(&User).Where("id = ?", userID).Updates(
		model.User{
			Email:    User.Email,
			Password: User.Password,
			Age:      User.Age,
			Username: User.Username,
		}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"age":        User.Age,
		"email":      User.Email,
		"password":   User.Password,
		"username":   User.Username,
		"updated_at": User.UpdatedAt,
	})
}

func UserDelete(c *gin.Context) {
	db := database.GetDB()
	User := model.User{}
	userID, _ := strconv.Atoi(c.Param("user_id"))

	db.Delete(&User, userID)
	db.Select("Photo").Delete(&User)

	c.JSON(http.StatusOK, gin.H{
		"message": "User has been deleted",
	})

}
