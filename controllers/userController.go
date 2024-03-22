package controllers

import (
	"MyGram/database"
	"MyGram/models"
	"MyGram/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm/clause"
	"net/http"
)

func UserRegister(c *gin.Context) {
	db := database.GetDB()

	User := models.User{}
	err := c.ShouldBind(&User)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(User)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	err = db.Create(&User).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"age":      User.Age,
		"email":    User.Email,
		"id":       User.ID,
		"username": User.Username,
	})
}

func UserLogin(c *gin.Context) {
	db := database.GetDB()

	request := models.User{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	User := models.User{}
	err = db.Debug().Where("email = ?", request.Email).Take(&User).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email/password",
		})
		return
	}

	comparePass := utils.ComparePass([]byte(User.Password), []byte(request.Password))
	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email/password",
		})
		return
	}

	token := utils.GenerateToken(User.ID, User.Email)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func UserUpdate(c *gin.Context) {
	db := database.GetDB()

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	User := models.User{}
	err := c.ShouldBind(&User)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	User.ID = userID
	err = db.Model(&User).Where("id = ?", userID).Updates(models.User{Email: User.Email, Username: User.Username}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, User)
}

func UserDelete(c *gin.Context) {
	db := database.GetDB()

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	User := models.User{}
	err := db.First(&User, userID).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Bad Request",
			"message": "User not found",
		})
		return
	}

	err = db.Delete(&User).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}

func UserMe(c *gin.Context) {
	db := database.GetDB()

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	User := models.User{}
	err := db.First(&User, userID).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Bad Request",
			"message": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, User)
}

func GetAllUsers(c *gin.Context) {
	db := database.GetDB()

	var Users []models.User
	db.Preload(clause.Associations).Find(&Users)
	c.JSON(http.StatusOK, Users)
}
