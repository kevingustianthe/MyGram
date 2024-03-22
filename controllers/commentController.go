package controllers

import (
	"MyGram/database"
	"MyGram/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func AddComment(c *gin.Context) {
	db := database.GetDB()

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	Comment := models.Comment{}
	err := c.ShouldBind(&Comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	Comment.UserID = userID
	err = db.Create(&Comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Comment)
}

func GetComments(c *gin.Context) {
	db := database.GetDB()

	var Comments []models.Comment
	err := db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, email, username")
	}).Preload("Photo", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, title, caption, photo_url, user_id")
	}).Find(&Comments).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Comments)
}

func UpdateComment(c *gin.Context) {
	db := database.GetDB()

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	commentID, _ := strconv.Atoi(c.Param("id"))

	Comment := models.Comment{}
	err := c.ShouldBind(&Comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	Comment.UserID = userID
	Comment.ID = uint(commentID)

	err = db.Model(&Comment).Where("id = ?", commentID).Updates(models.Comment{
		Message: Comment.Message,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Comment)
}

func DeleteComment(c *gin.Context) {
	db := database.GetDB()

	Comment := models.Comment{}
	commentID, _ := strconv.Atoi(c.Param("id"))

	err := db.First(&Comment, commentID).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Bad Request",
			"message": "Comment not found",
		})
		return
	}

	err = db.Delete(&Comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
