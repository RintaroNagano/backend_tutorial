package controller

import (
	"net/http"

	"crypto/sha256"
	"encoding/hex"

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

	checkSum := sha256.Sum256([]byte(password))
	hashpass := hex.EncodeToString(checkSum[:])

	user := model.User{}

	db.GetDB().Where("user_id = ?", userId).First(&user)
	// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;

	if (user == model.User{}) {
		db.GetDB().Create(&model.User{UserId: userId, Password: hashpass})
		c.JSON(http.StatusCreated, gin.H{
			"token": "hogefugapiyo",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "This user ID has been used",
		})
	}
}

func SigninHandler(c *gin.Context) {
	userId := c.PostForm("UserId")
	password := c.PostForm("Password")

	checkSum := sha256.Sum256([]byte(password))
	hashpass := hex.EncodeToString(checkSum[:])

	var user model.User
	// Read
	db.GetDB().First(&user, "user_id = ?", userId)

	if hashpass == user.Password {
		c.JSON(http.StatusOK, gin.H{
			"token": "hogefugapiyo",
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid password",
		})
	}
}
