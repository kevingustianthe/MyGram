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

func AddPhoto(c *gin.Context) {
	db := database.GetDB()

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	Photo := models.Photo{}
	err := c.ShouldBind(&Photo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	Photo.UserID = userID
	err = db.Create(&Photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Photo)
}

func GetPhotos(c *gin.Context) {
	db := database.GetDB()

	var Photos []models.Photo
	err := db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, email, username")
	}).Find(&Photos).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Photos)
}

func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	photoID, _ := strconv.Atoi(c.Param("id"))

	Photo := models.Photo{}
	err := c.ShouldBind(&Photo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	Photo.UserID = userID
	Photo.ID = uint(photoID)

	err = db.Model(&Photo).Where("id = ?", photoID).Updates(models.Photo{
		Title:    Photo.Title,
		Caption:  Photo.Caption,
		PhotoURL: Photo.PhotoURL,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Photo)
}

func DeletePhoto(c *gin.Context) {
	db := database.GetDB()

	Photo := models.Photo{}
	photoID, _ := strconv.Atoi(c.Param("id"))

	err := db.First(&Photo, photoID).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Bad Request",
			"message": "Photo not found",
		})
		return
	}

	err = db.Delete(&Photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
