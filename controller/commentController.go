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

func CreateComment(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)

	//Get ID User yang sedang login
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	Comment := model.Comment{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID

	err := db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoID,
		"user_id":    Comment.UserID,
		"created_at": Comment.CreatedAt,
	})
}

func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	Comment := model.Comment{}
	commentID, _ := strconv.Atoi(c.Param("comment_id"))

	err := db.Delete(&Comment, commentID).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Comment has been deleted",
	})
}

func GetMyComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	Comment := []model.Comment{}

	err := db.Where("user_id = ?", userID).Find(&Comment)
	if err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error,
		})
		return
	}

	MyComment := []model.MyComment{}

	for i := 0; i < len(Comment); i++ {
		tempMyComment := model.MyComment{}
		tempPhoto := model.Photo{}

		_ = db.Find(&tempPhoto, Comment[i].PhotoID)

		tempMyComment.ID = Comment[i].ID
		tempMyComment.Message = Comment[i].Message
		tempMyComment.PhotoID = Comment[i].PhotoID
		tempMyComment.UpdatedAt = Comment[i].UpdatedAt
		tempMyComment.CreatedAt = Comment[i].CreatedAt

		tempMyComment.Photo.ID = tempPhoto.ID
		tempMyComment.Photo.Title = tempPhoto.Title
		tempMyComment.Photo.Caption = tempPhoto.Caption
		tempMyComment.Photo.Photo_url = tempPhoto.Photo_url
		tempMyComment.Photo.UserID = tempPhoto.UserID

		MyComment = append(MyComment, tempMyComment)
	}

	c.JSON(http.StatusOK, MyComment)

}

func EditComment(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)
	commentID, _ := strconv.Atoi(c.Param("comment_id"))
	Comment := model.Comment{}
	EditedComment := model.Comment{}
	_ = db.First(&Comment, commentID)

	if contentType == appJSON {
		c.ShouldBindJSON(&EditedComment)
	} else {
		c.ShouldBind(&EditedComment)
	}

	EditedComment.ID = Comment.ID
	EditedComment.UserID = Comment.UserID
	EditedComment.PhotoID = Comment.PhotoID

	err := db.Model(&EditedComment).Where("id=?", commentID).Updates(
		model.Comment{
			Message: EditedComment.Message,
		}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         EditedComment.ID,
		"message":    EditedComment.Message,
		"photo_id":   EditedComment.PhotoID,
		"user_id":    EditedComment.UserID,
		"updated_at": EditedComment.UpdatedAt,
	})

}
