package controllers

import (
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/mhdiiilham/ginorm/db"
	h "github.com/mhdiiilham/ginorm/helpers"
	m "github.com/mhdiiilham/ginorm/models"
)

// CreateUser ...
func CreateUser(c *gin.Context) {
	// Declare varible that needed
	var body m.UserSignUpInput

	// Validate user's input
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"errors": err.Error()})
		return
	}

	// Create user
	user := m.User{
		FirstName:      body.FirstName,
		LastName:       body.LastName,
		Email:          body.Email,
		PasswordHashed: h.HashPassword([]byte(body.Password)),
	}

	// Save it to DB if there's no error
	if err := db.MySQL().Save(&user); err != nil {
		c.JSON(500, gin.H{"message": "Internal Server Error. Failed create your account, please try again"})
		return
	}

	// Create access token
	token, err := h.CreateJWTToken(user.ID, user.Email)

	if err != nil {
		log.Warn("Error at creating token after signup. Error: ", err.Error())
		c.JSON(500, gin.H{"errors": "Internal Server Error"})
		return
	}

	// Send token to client
	c.JSON(201, gin.H{
		"message": "Sign Up Success",
		"email":   user.Email,
		"token":   token,
	})

}

// Login ...
func Login(c *gin.Context) {
	// Declare variables that needed
	var body m.UserLoginInput
	var user m.User

	// Validate user's Input
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"errors": err.Error()})
		return
	}

	// Find user with given email
	// and assign it to `user`
	find := db.MySQL().Where("email = ?", body.Email).Find(&user)

	if find.Error != nil {
		c.JSON(400, gin.H{"errors": "Email / Password is Wrong!"})
		return
	}

	// Compare user's password
	if isValid := h.ComparePassword(user.PasswordHashed, []byte(body.Password)); !isValid {
		c.JSON(400, gin.H{"errors": "Email / Password is Wrong!"})
		return
	}

	// If password is right
	// Create access token
	token, err := h.CreateJWTToken(user.ID, user.Email)
	if err != nil {
		log.Warn("Error at creating token after signup. Error: ", err.Error())
		c.JSON(500, gin.H{"errors": "Internal Server Error"})
		return
	}

	// Send token to client
	c.JSON(200, gin.H{
		"message": "Sign In Success",
		"email":   user.Email,
		"token":   token,
	})

}
