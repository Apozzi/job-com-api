package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
)

type UserLogin struct {
	Username string
	Password string
}

func Login(c *gin.Context) {
	var userLogin UserLogin
	c.BindJSON(&userLogin)
	var user User
	db.Where("email = ? AND password = ?", userLogin.Username, userLogin.Password).Find(&user)
	if user.IdUser != 0 {
		token, _ := auth.CreateToken(user.Email, time.Hour*12)
		c.JSON(200, gin.H{
			"token":   token,
			"user_id": user.IdUser,
		})
	}
}

func VerifyToken(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": true,
	})
}
