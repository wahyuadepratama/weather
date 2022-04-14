package controllers

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wahyuadepratama/weather/src/config"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

// Define database client
var db *gorm.DB = config.ConnectDB()

func GenereateToken(data string) string {

	mySigningKey := []byte(os.Getenv("SECRET_KEY"))

	type MyCustomClaims struct {
		Email string `json:"email"`
		jwt.StandardClaims
	}

	sessionDay, _ := strconv.Atoi(os.Getenv("SESSION_DAY"))

	// Create the Claims
	claims := MyCustomClaims{
		data,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * time.Duration(sessionDay)).Unix(),
			Issuer:    data,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return "Token Invalid"
	}

	return ss
}

func VerifyToken(tokenString string) (int, string) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return 0, "Token invalid"
	}

	if token.Valid {
		return 1, "Token valid"
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return 0, "That's not even a token"
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			return 0, "Your token is expired"
		} else {
			return 0, "We cannot handle this error"
		}
	} else {
		return 0, "We cannot handle this error"
	}
}

func IsTokenValid(context *gin.Context) int {
	type DataRequest struct {
		Token string `json:"token"`
	}

	var request DataRequest

	// Binding request body json to request body struct
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  "401",
			"message": "Token required",
		})
		return 0
	}

	response, data := VerifyToken(request.Token)
	if response != 1 {
		context.JSON(http.StatusOK, gin.H{
			"status":  "401",
			"message": data,
		})
		return 0
	}

	return 1
}
