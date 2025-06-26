package controller

import (
	"fmt" // 🔑 Добавь этот импорт
	"github.com/gin-gonic/gin"
	"hotel-booking/internal/errs"
	"hotel-booking/internal/models"
	"hotel-booking/internal/service"
	"hotel-booking/internal/utils"
	"net/http"
)

func SignUp(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := service.CreateUser(user); err != nil {
		if err == errs.ErrUserAlreadyExists {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

func SignIn(c *gin.Context) {
	var input models.UserSignIn
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// ⚡ Можно также добавить отладку для входа:
	fmt.Printf("DEBUG SignIn: Username=%s Password=%s\n", input.Username, input.Password)

	user, err := service.GetUserByUsernameAndPassword(input.Username, input.Password)
	if err != nil {
		if err == errs.ErrIncorrectUsernameOrPassword {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
