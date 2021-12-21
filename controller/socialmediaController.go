package controller

import (
	"MyGram/database"
	"MyGram/helper"
	"MyGram/model"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateSocialmedia(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	Socialmedia := model.SocialMedia{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Socialmedia)
	} else {
		c.ShouldBind(&Socialmedia)
	}

	Socialmedia.UserID = userID

	err := db.Debug().Create(&Socialmedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":               Socialmedia.ID,
		"name":             Socialmedia.Name,
		"social_media_url": Socialmedia.SocialMediaURL,
		"user_id":          Socialmedia.UserID,
		"created_at":       Socialmedia.CreatedAt,
	})

}

func GetMySocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	Socialmedia := []model.SocialMedia{}

	err := db.Where("user_id = ?", userID).Find(&Socialmedia)
	if err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error,
		})
		return
	}

	MySocialmedia := []model.MySocialMedia{}

	for i := 0; i < len(Socialmedia); i++ {
		tempMySocialmedia := model.MySocialMedia{}
		tempUser := model.User{}

		_ = db.Find(&tempUser, Socialmedia[i].UserID)

		tempMySocialmedia.ID = Socialmedia[i].ID
		tempMySocialmedia.Name = Socialmedia[i].Name
		tempMySocialmedia.SocialMediaURL = Socialmedia[i].SocialMediaURL
		tempMySocialmedia.UserID = Socialmedia[i].UserID
		tempMySocialmedia.CreatedAt = Socialmedia[i].CreatedAt
		tempMySocialmedia.User.ID = tempUser.ID
		tempMySocialmedia.User.Username = tempUser.Username
		tempMySocialmedia.User.Email = tempUser.Email

		MySocialmedia = append(MySocialmedia, tempMySocialmedia)
	}

	c.JSON(http.StatusOK, MySocialmedia)
}

func EditSocialMedia(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)
	socialmediaID, _ := strconv.Atoi(c.Param("socialmedia_id"))
	Socialmedia := model.SocialMedia{}
	EditedSocialmedia := model.SocialMedia{}
	_ = db.First(&Socialmedia, socialmediaID)

	if contentType == appJSON {
		c.ShouldBindJSON(&EditedSocialmedia)
	} else {
		c.ShouldBind(&EditedSocialmedia)
	}

	EditedSocialmedia.ID = Socialmedia.ID
	EditedSocialmedia.UserID = Socialmedia.UserID

	err := db.Model(&EditedSocialmedia).Where("id=?", socialmediaID).Updates(
		model.SocialMedia{
			Name:           EditedSocialmedia.Name,
			SocialMediaURL: EditedSocialmedia.SocialMediaURL,
		}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":               EditedSocialmedia.ID,
		"name":             EditedSocialmedia.Name,
		"social_media_url": EditedSocialmedia.SocialMediaURL,
		"user_id":          EditedSocialmedia.UserID,
		"updated_at":       EditedSocialmedia.UpdatedAt,
	})

}

func DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()
	SocialMedia := model.SocialMedia{}
	socialmediaID, _ := strconv.Atoi(c.Param("socialmedia_id"))

	err := db.Delete(&SocialMedia, socialmediaID).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Social Media has been deleted",
	})

}
