package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hotel-booking/internal/errs"
	"hotel-booking/internal/models"
	"hotel-booking/internal/service"
	"hotel-booking/internal/utils"
	"net/http"
)

// SignUp godoc
// @Summary Регистрация пользователя
// @Description Регистрирует нового пользователя
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.User true "Данные пользователя"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/sign-up [post]
func SignUp(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	user.Role = "user" // по умолчанию роль user
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

// SignIn godoc
// @Summary Вход пользователя
// @Description Аутентификация пользователя и выдача JWT
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body models.UserSignIn true "Данные для входа"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/sign-in [post]
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

	token, err := utils.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
