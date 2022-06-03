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
	// SELECT * FROM users WHERE user_id = '(valuable userId)' ORDER BY id LIMIT 1;

	if (user == model.User{}) {
		db.GetDB().Create(&model.User{UserId: userId, Password: hashpass})

		token, err := generateToken(userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to generate token",
			})
			return
		}

		c.SetCookie("jwt", token, 300, "/", "localhost", false, true)
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

	if hashpass != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid password",
		})
		return
	}

	jwtString, err := c.Cookie("jwt")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "You are not authorized",
		})
		return
	}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	claims := &model.JwtClaims{}
	result, err := jwt.ParseWithClaims(jwtString, claims, func(token *jwt.Token) (interface{}, error) {
		return constants.Get_const_JWT_KEY(), nil
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to parse token",
		})
		return
	}
	if claims.UserId != userId {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "You use invalid token",
		})
		return
	}
	if !result.Valid {
		if checkTokenExpiration(claims) {
			// if the jwt expires, regenerate and reset cookie
			jwtString, err = generateToken(userId)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Failed to generate token",
				})
			}
			c.SetCookie("jwt", jwtString, 300, "/", "localhost", false, true)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "You are not authorized",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"token": jwtString,
	})
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

func checkTokenExpiration(claims *model.JwtClaims) bool {
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		return true
	}
	return false
}
