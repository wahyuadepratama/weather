package controllers

import (
	"net/http"
	"net/mail"
	"strconv"
	"time"

	"github.com/wahyuadepratama/weather/src/models"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

func UserLogin(context *gin.Context) {

	type DataLogin struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var dataLogin DataLogin

	// Binding request body json to request body struct
	if err := context.ShouldBindJSON(&dataLogin); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  "401",
		})
		return
	}

	// Validation input data
	if dataLogin.Email == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "401",
			"message": "Login failed, email can't be empty",
		})
		return
	}

	if dataLogin.Password == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "401",
			"message": "Login failed, password can't be empty",
		})
		return
	}

	// Check is user data exists
	var user models.User
	usersData := db.Raw("select * from users where email = ? limit 1", dataLogin.Email).Scan(&user)
	if usersData.Error != nil || user.Email == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "401",
			"message": "Login failed, email didn't registered yet",
		})
		return
	}

	// Check is password match or not
	errPassCheck := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dataLogin.Password))
	if errPassCheck != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "401",
			"message": "Login failed, wrong password",
		})
		return
	}

	// Generate token
	token := GenereateToken(dataLogin.Email)

	// Creating http response
	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Login Success",
		"data": gin.H{
			"email": user.Email,
			"name":  user.Name,
			"token": token,
		},
	})

}

func UserRegister(context *gin.Context) {

	// Define struct
	type DataRegis struct {
		Email    string `json:"email"`
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	var dataRegis DataRegis

	// Binding request body json to request body struct
	if err := context.ShouldBindJSON(&dataRegis); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  "401",
		})
		return
	}

	// Validation input data
	if dataRegis.Email == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "401",
			"message": "Registration failed, email can't be empty",
		})
		return
	}

	if dataRegis.Name == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "401",
			"message": "Registration failed, name can't be empty",
		})
		return
	}

	if dataRegis.Password == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "401",
			"message": "Registration failed, password can't be empty",
		})
		return
	}

	_, err := mail.ParseAddress(dataRegis.Email)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "401",
			"message": "Registration failed, email format is not valid",
		})
		return
	}

	// Generate UD & Hash password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(dataRegis.Password), 8)
	ID := "U" + strconv.FormatInt(time.Now().Unix(), 10)
	token := GenereateToken(dataRegis.Email)

	// Insert data to users table
	errInsertUser := db.Exec("INSERT INTO users (id,email,name,password,created_at) VALUES (?,?,?,?,now())", ID, dataRegis.Email, dataRegis.Name, string(hashedPassword))

	if errInsertUser.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "401",
			"message": "Registration failed",
		})
		return
	}

	// Creating http response
	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Registration Success",
		"data": gin.H{
			"email": dataRegis.Email,
			"name":  dataRegis.Name,
			"token": token,
		},
	})
}
