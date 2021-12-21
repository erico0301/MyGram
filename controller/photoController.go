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

func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)

	//Get ID User yang sedang login
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	Photo := model.Photo{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}
	Photo.UserID = userID

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.Photo_url,
		"user_id":    Photo.UserID,
		"created_at": Photo.CreatedAt,
	})

}

func PhotoDelete(c *gin.Context) {
	db := database.GetDB()
	Photo := model.Photo{}
	photoID, _ := strconv.Atoi(c.Param("photo_id"))

	db.Delete(&Photo, photoID)

	c.JSON(http.StatusOK, gin.H{
		"message": "Photo has been deleted",
	})
}

func GetMyPhoto(c *gin.Context) {
	db := database.GetDB()
	//Get ID User yang sedang login
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	Photo := []model.Photo{}

	err := db.Where("user_id = ?", strconv.Itoa(int(userID))).Find(&Photo)

	if err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error,
		})
		return
	}

	MyPhoto := []model.MyPhoto{}

	for i := 0; i < len(Photo); i++ {
		photoID := Photo[i].ID
		comment := []model.Comment{}
		tempMyPhoto := model.MyPhoto{}

		err := db.Where("photo_id=?", photoID).Find(&comment)
		if err.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error,
			})
			return
		}
		tempMyPhoto.ID = Photo[i].ID
		tempMyPhoto.Title = Photo[i].Title
		tempMyPhoto.Caption = Photo[i].Caption
		tempMyPhoto.Photo_url = Photo[i].Photo_url
		tempMyPhoto.CreatedAt = Photo[i].CreatedAt
		if len(comment) > 0 {
			for j := 0; j < len(comment); j++ {
				tempComment := model.GetPhotoComment{}
				tempComment.ID = comment[j].ID
				tempComment.Message = comment[j].Message
				tempComment.UserID = comment[j].UserID
				tempComment.CreatedAt = comment[j].CreatedAt
				tempMyPhoto.Comments = append(tempMyPhoto.Comments, tempComment)
			}
		}
		MyPhoto = append(MyPhoto, tempMyPhoto)
	}

	c.JSON(http.StatusOK, MyPhoto)

}

func GetPhotoByID(c *gin.Context) {
	db := database.GetDB()
	tempPhoto := model.PublicPhoto{}
	Photo := model.Photo{}
	photoID, _ := strconv.Atoi(c.Param("photo_id"))

	err := db.Find(&Photo, photoID)

	if err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error,
		})
		return
	}

	if Photo.ID == 0 || Photo.UserID < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "PhotoID not found",
		})
		return
	}

	tempPhoto.ID = Photo.ID
	tempPhoto.Title = Photo.Title
	tempPhoto.Caption = Photo.Caption
	tempPhoto.Photo_url = Photo.Photo_url
	tempPhoto.CreatedAt = Photo.CreatedAt

	User := model.User{}
	_ = db.Find(&User, Photo.UserID)

	tempPhoto.User.ID = User.ID
	tempPhoto.User.Username = User.Username
	tempPhoto.User.Email = User.Email

	comment := []model.Comment{}
	_ = db.Where("photo_id=?", photoID).Find(&comment)
	if len(comment) > 0 {
		for j := 0; j < len(comment); j++ {
			tempComment := model.GetPhotoComment{}
			tempComment.ID = comment[j].ID
			tempComment.Message = comment[j].Message
			tempComment.UserID = comment[j].UserID
			tempComment.CreatedAt = comment[j].CreatedAt
			tempPhoto.Comments = append(tempPhoto.Comments, tempComment)
		}
	}

	c.JSON(http.StatusOK, tempPhoto)
}

func EditPhoto(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)
	Photo := model.Photo{}
	photoID, _ := strconv.Atoi(c.Param("photo_id"))
	MyPhoto := model.Photo{}
	_ = db.First(&MyPhoto, photoID)

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.ID = MyPhoto.ID
	Photo.UserID = MyPhoto.UserID

	err := db.Model(&Photo).Where("id=?", photoID).Updates(
		model.Photo{
			Title:     Photo.Title,
			Caption:   Photo.Caption,
			Photo_url: Photo.Photo_url,
		}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.Photo_url,
		"user_id":    Photo.UserID,
		"updated_at": Photo.UpdatedAt,
	})

}
