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

func AddSocialMedia(c *gin.Context) {
	db := database.GetDB()

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	SocialMedia := models.SocialMedia{}
	err := c.ShouldBind(&SocialMedia)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	SocialMedia.UserID = userID
	err = db.Create(&SocialMedia).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, SocialMedia)
}

func GetSocialMedias(c *gin.Context) {
	db := database.GetDB()

	var SocialMedias []models.SocialMedia
	err := db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, email, username")
	}).Find(&SocialMedias).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SocialMedias)
}

func UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	socialMediaID, _ := strconv.Atoi(c.Param("id"))

	SocialMedia := models.SocialMedia{}
	err := c.ShouldBind(&SocialMedia)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	SocialMedia.UserID = userID
	SocialMedia.ID = uint(socialMediaID)

	err = db.Model(&SocialMedia).Where("id = ?", socialMediaID).Updates(models.SocialMedia{
		Name:           SocialMedia.Name,
		SocialMediaURL: SocialMedia.SocialMediaURL,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SocialMedia)
}

func DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()

	SocialMedia := models.SocialMedia{}
	socialMediaID, _ := strconv.Atoi(c.Param("id"))

	err := db.First(&SocialMedia, socialMediaID).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Bad Request",
			"message": "Social media not found",
		})
		return
	}

	err = db.Delete(&SocialMedia).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}
