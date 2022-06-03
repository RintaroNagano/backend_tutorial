package controller

import (
	"net/http"
	"time"

	"crypto/sha256"
	"encoding/hex"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"sample/constants"
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

		token, err := generateToken(userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to generate token",
			})
		}

		c.JSON(http.StatusCreated, gin.H{
			"token": token,
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

func generateToken(userId string) (string, error) {
	expirationTime := time.Now().Add(constants.EXPIRATION_TIME)

	claims := &model.JwtClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString(constants.Get_const_JWT_KEY())
	if err != nil {
		return tokenString, err
	}

	return tokenString, err
}
