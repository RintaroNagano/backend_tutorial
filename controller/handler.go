package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"sample/db"
	"sample/model"
)

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ping",
	})
}

func SignupHandler(c *gin.Context) {
	userId := c.PostForm("UserId")
	password := c.PostForm("Password")

	intUserId, _ := strconv.Atoi(userId)
	// Create
	db.DB.Create(&model.User{UserId: intUserId, Password: password, Token: "hogefugapiyo"})

	c.JSON(http.StatusCreated, gin.H{
		"token": "hogefugapiyo",
	})
}

func SigninHandler(c *gin.Context) {
	userId := c.PostForm("UserId")
	password := c.PostForm("Password")

	var user model.User
	// Read
	db.DB.First(&user, "user_id = ?", userId)

	if password == user.Password {
		c.JSON(http.StatusOK, gin.H{
			"token": user.Token,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid password",
		})
	}
}
